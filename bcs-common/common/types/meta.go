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

package types

import "time"

//BcsDataType data type for bcs-api & client
type BcsDataType string

const (
	BcsDataType_APP              BcsDataType = "application"
	BcsDataType_PROCESS          BcsDataType = "process"
	BcsDataType_POD              BcsDataType = "taskgroup"
	BcsDataType_SECRET           BcsDataType = "secret"
	BcsDataType_CONFIGMAP        BcsDataType = "configmap"
	BcsDataType_SERVICE          BcsDataType = "service"
	BcsDataType_LOADBALANCE      BcsDataType = "loadbalance"
	BcsDataType_HAPROXY          BcsDataType = "haproxy"
	BcsDataType_EXPORTSERVICE    BcsDataType = "exportservice"
	BcsDataType_DEPLOYMENT       BcsDataType = "deployment"
	BcsDataType_CRR              BcsDataType = "crr"
	BcsDataType_WebConsole       BcsDataType = "webconsole"
	BcsDataType_Admissionwebhook BcsDataType = "admissionwebhook"
)

//TypeMeta for bcs data type
type TypeMeta struct {
	APIVersion string      `json:"apiVersion"`
	Kind       BcsDataType `json:"kind"`
}

//ObjectMeta for meta for datatype
type ObjectMeta struct {
	Name              string            `json:"name"`
	NameSpace         string            `json:"namespace"`
	CreationTimestamp time.Time         `json:"creationTimestamp,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	ClusterName       string            `json:"clusterName,omitempty"`
}

//GetName for ObjectMeta
func (om *ObjectMeta) GetName() string {
	return om.Name
}

//SetName set object name
func (om *ObjectMeta) SetName(name string) {
	om.Name = name
}

//GetNamespace for ObjectMeta
func (om *ObjectMeta) GetNamespace() string {
	return om.NameSpace
}

//SetNamespace set object namespace
func (om *ObjectMeta) SetNamespace(ns string) {
	om.NameSpace = ns
}

//GetCreationTimestamp get create timestamp
func (om *ObjectMeta) GetCreationTimestamp() time.Time {
	return om.CreationTimestamp
}

//SetCreationTimestamp set creat timestamp
func (om *ObjectMeta) SetCreationTimestamp(timestamp time.Time) {
	om.CreationTimestamp = timestamp
}

//GetLabels for ObjectMeta
func (om *ObjectMeta) GetLabels() map[string]string {
	return om.Labels
}

//SetLabels set objec labels
func (om *ObjectMeta) SetLabels(labels map[string]string) {
	om.Labels = labels
}

//GetAnnotations for ObjectMeta
func (om *ObjectMeta) GetAnnotations() map[string]string {
	return om.Annotations
}

//SetAnnotations get annotation name
func (om *ObjectMeta) SetAnnotations(annotation map[string]string) {
	om.Annotations = annotation
}

//GetClusterName get cluster name
func (om *ObjectMeta) GetClusterName() string {
	return om.ClusterName
}

//SetClusterName set cluster name
func (om *ObjectMeta) SetClusterName(clusterName string) {
	om.ClusterName = clusterName
}

//KeyToValue key/value structs
type KeyToValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
