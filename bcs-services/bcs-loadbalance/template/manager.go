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

package template

import (
	"bk-bcs/bcs-common/common/metric"
	"bk-bcs/bcs-services/bcs-loadbalance/types"
	"os"
)

var (
	// HealthStatusOK healthy flag
	HealthStatusOK = true
	// HealthStatusOKMsg messsage when healthy
	HealthStatusOKMsg = "I am OK"
	// HealthStatusNotOK unhealthy flag
	HealthStatusNotOK = false
)

// IsFileExist check file exist
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// Manager define interface for manager
type Manager interface {
	//Start point, do not block
	Start() error
	//Stop
	Stop()
	//Create config file with tmpData,
	Create(tmpData *types.TemplateData) (string, error)
	//CheckDifference two file are difference, true is difference
	CheckDifference(oldFile, curFile string) bool
	//Validate new cfg file grammar is OK
	Validate(newFile string) bool
	//Replace old cfg file with cur one, return old file backup
	Replace(oldFile, curFile string) error
	//Reload haproxy with new config file
	Reload(cfgFile string) error
	//GetHealthInfo response healthz info
	GetHealthInfo() metric.HealthMeta
	//Get metric meta
	GetMetricMeta() *metric.MetricMeta
	//Get metric result
	GetMetricResult() (*metric.MetricResult, error)
}
