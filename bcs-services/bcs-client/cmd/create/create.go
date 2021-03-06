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

package create

import (
	"fmt"

	"bk-bcs/bcs-services/bcs-client/cmd/utils"

	"github.com/urfave/cli"
)

func NewCreateCommand() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "create new application/process/service/secret/configmap/deployment",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "from-file, f",
				Usage: "Create with configuration `FILE`",
			},
			cli.StringFlag{
				Name:  "clusterid",
				Usage: "Cluster ID",
			},
			cli.StringFlag{
				Name:  "type, t",
				Usage: "Create type, value can be app/service/secret/configmap/deployment",
			},
		},
		Action: func(c *cli.Context) error {
			if err := create(utils.NewClientContext(c)); err != nil {
				return err
			}
			return nil
		},
	}
}

func create(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionType); err != nil {
		return err
	}

	resourceType := c.String(utils.OptionType)

	switch resourceType {
	case "app", "application":
		return createApplication(c)
	case "process":
		return createProcess(c)
	case "configmap":
		return createConfigMap(c)
	case "secret":
		return createSecret(c)
	case "service":
		return createService(c)
	case "deploy", "deployment":
		return createDeployment(c)
	default:
		return fmt.Errorf("invalid type: %s", resourceType)
	}
}
