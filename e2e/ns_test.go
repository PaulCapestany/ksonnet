// Copyright 2018 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// +build e2e

package e2e

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("ks ns", func() {
	var a app

	BeforeEach(func() {
		a = e.initApp("")
		a.generateDeployedService()
	})

	Describe("create", func() {
		Context("it does not exist", func() {
			It("creates the namespace", func() {
				o := a.runKs("ns", "create", "new-ns")
				assertExitStatus(o, 0)
			})
		})
		Context("it exists", func() {
			It("returns an error", func() {
				o := a.runKs("ns", "create", "/")
				assertExitStatus(o, 1)
				assertOutput("ns/create/invalid.txt", o.stderr)
			})
		})
	})

	Describe("list", func() {
		It("shows a list of namespaces", func() {
			o := a.runKs("ns", "list")
			assertExitStatus(o, 0)
			assertOutput("ns/list/output.txt", o.stdout)
		})

		Context("with a valid environment", func() {
			It("shows a list of namespaces", func() {
				o := a.runKs("ns", "list", "--env", "default")
				assertExitStatus(o, 0)
				assertOutput("ns/list/output.txt", o.stdout)
			})
		})

		Context("with an invalid environment", func() {
			It("returns an error", func() {
				o := a.runKs("ns", "list", "--env", "invalid")
				assertExitStatus(o, 1)
				assertOutput("ns/list/invalid.txt", o.stderr)
			})
		})
	})
})
