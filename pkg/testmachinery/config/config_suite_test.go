// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/gardener/test-infra/pkg/testmachinery/config"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config", func() {

	Context("Create", func() {

		DescribeTable("create config elements",
			func(configs []tmv1beta1.ConfigElement) {
				newConfigs := config.New(configs, config.LevelTestDefinition)
				Expect(len(configs)).To(Equal(len(newConfigs)))
				for i, newElem := range newConfigs {
					Expect(*newElem.Info).To(Equal(configs[i]))
					Expect(newElem.Name).ToNot(Equal(""))
				}
			},
			Entry("1 element", []tmv1beta1.ConfigElement{{Name: "test1", Value: "", Type: "env"}}),
			Entry("3 elements", []tmv1beta1.ConfigElement{{Name: "test1", Value: "", Type: "env"},
				{Name: "test2", Value: "", Type: "env"}, {Name: "test3", Value: "", Type: "env"}}),
			Entry("0 elements", []tmv1beta1.ConfigElement{}),
		)

	})

})
