// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"fmt"

	"github.com/kubeflow/pipelines/backend/src/common/util"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

type WorkflowClientFake struct {
	workflows map[string]util.ExecutionSpec
}

func NewWorkflowClientFake() *WorkflowClientFake {
	return &WorkflowClientFake{
		workflows: make(map[string]util.ExecutionSpec),
	}
}

func (p *WorkflowClientFake) Get(namespace string, name string) (
	util.ExecutionSpec, error) {
	workflow, ok := p.workflows[getKey(namespace, name)]
	if !ok {
		return nil, util.NewCustomError(fmt.Errorf("Error"),
			util.CUSTOM_CODE_NOT_FOUND, "Workflow not found: %s/%s", namespace, name)
	}
	if workflow == nil {
		return nil, util.NewCustomError(fmt.Errorf("Error"), util.CUSTOM_CODE_GENERIC,
			"Error getting workflow: %s/%s", namespace, name)
	}
	return workflow, nil
}

func (p *WorkflowClientFake) Put(namespace string, name string, wf util.ExecutionSpec) {
	p.workflows[getKey(namespace, name)] = wf
}

func getKey(namespace string, name string) string {
	return namespace + "/" + name
}
