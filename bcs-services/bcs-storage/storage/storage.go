/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package bcsstorage

import (
	"net/http"
	"net/http/pprof"

	"bk-bcs/bcs-common/common/blog"
	"bk-bcs/bcs-common/common/http/httpserver"
	"bk-bcs/bcs-common/common/metric"
	"bk-bcs/bcs-common/common/types"
	"bk-bcs/bcs-services/bcs-storage/app/options"
	"bk-bcs/bcs-services/bcs-storage/storage/actions"
	"bk-bcs/bcs-services/bcs-storage/storage/apiserver"
	"bk-bcs/bcs-services/bcs-storage/storage/rdiscover"

	"github.com/emicklei/go-restful"
)

// StorageServer is a data struct of bcs storage server
type StorageServer struct {
	conf       *options.StorageOptions
	httpServer *httpserver.HttpServer
	rd         *rdiscover.RegDiscover
}

// NewStorageServer create storage server object
func NewStorageServer(op *options.StorageOptions) (*StorageServer, error) {
	s := &StorageServer{}

	// Configuration
	s.conf = op

	// Http server
	s.httpServer = httpserver.NewHttpServer(s.conf.Port, s.conf.Address, "")
	if s.conf.ServerCert.IsSSL {
		s.httpServer.SetSsl(s.conf.ServerCert.CAFile, s.conf.ServerCert.CertFile, s.conf.ServerCert.KeyFile, s.conf.ServerCert.CertPwd)
	}

	// RDiscover
	s.rd = rdiscover.NewRegDiscover(s.conf)

	// ApiResource
	a := apiserver.GetAPIResource()
	a.SetConfig(op)
	a.InitActions()

	return s, nil
}

func (s *StorageServer) initHTTPServer() error {
	a := apiserver.GetAPIResource()

	// Api v1
	s.httpServer.RegisterWebServer(actions.PathV1, nil, a.ActionsV1)

	if a.Conf.DebugMode {
		s.initDebug()
	}
	return nil
}

func (s *StorageServer) initDebug() {
	action := []*httpserver.Action{
		httpserver.NewAction("GET", "/debug/pprof/", nil, getRouteFunc(pprof.Index)),
		httpserver.NewAction("GET", "/debug/pprof/{uri:*}", nil, getRouteFunc(pprof.Index)),
		httpserver.NewAction("GET", "/debug/pprof/cmdline", nil, getRouteFunc(pprof.Cmdline)),
		httpserver.NewAction("GET", "/debug/pprof/profile", nil, getRouteFunc(pprof.Profile)),
		httpserver.NewAction("GET", "/debug/pprof/symbol", nil, getRouteFunc(pprof.Symbol)),
		httpserver.NewAction("GET", "/debug/pprof/trace", nil, getRouteFunc(pprof.Trace)),
	}
	s.httpServer.RegisterWebServer("", nil, action)
}

// Start to run storage server
func (s *StorageServer) Start() error {
	chErr := make(chan error, 1)

	s.initHTTPServer()

	go func() {
		err := s.httpServer.ListenAndServe()
		blog.Errorf("http listen and service failed! err:%s", err.Error())
		chErr <- err
	}()

	// register and discover
	go func() {
		err := s.rd.Start()
		blog.Errorf("storage rdiscover start failed! err:%s", err.Error())
		chErr <- err
	}()

	metricHandler(s.conf)

	// startDaemon
	actions.StartActionDaemon()

	select {
	case err := <-chErr:
		blog.Errorf("exit! err:%s", err.Error())
		return err
	}
}

func metricHandler(op *options.StorageOptions) {
	c := metric.Config{
		ModuleName: types.BCS_MODULE_STORAGE,
		MetricPort: op.MetricPort,
		IP:         op.Address,
		RunMode:    metric.Master_Master_Mode,

		SvrCaFile:   op.ServerCert.CAFile,
		SvrCertFile: op.ServerCert.CertFile,
		SvrKeyFile:  op.ServerCert.KeyFile,
		SvrKeyPwd:   op.ServerCert.CertPwd,
	}

	if err := metric.NewMetricController(
		c,
		apiserver.GetHealth,
	); err != nil {
		blog.Errorf("metric server error: %v", err)
		return
	}
	blog.Infof("start metric server successfully")
}

func getRouteFunc(f http.HandlerFunc) restful.RouteFunction {
	return restful.RouteFunction(func(req *restful.Request, resp *restful.Response) {
		f(resp, req.Request)
	})
}
