// Copyright 2018 The ksonnet authors
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

package actions

import (
	"net/url"
	"strings"

	"github.com/ksonnet/ksonnet/metadata/app"
	"github.com/ksonnet/ksonnet/pkg/registry"
)

// RunRegistryAdd runs `registry add`
func RunRegistryAdd(ksApp app.App, name, uri, version string, isOverride bool) error {
	ra, err := NewRegistryAdd(ksApp, name, uri, version, isOverride)
	if err != nil {
		return err
	}

	return ra.Run()
}

// RegistryAdd lists namespaces.
type RegistryAdd struct {
	app         app.App
	name        string
	uri         string
	version     string
	isOverride  bool
	registryAdd func(a app.App, name, protocol, uri, version string, isOverride bool) (*registry.Spec, error)
}

// NewRegistryAdd creates an instance of RegistryAdd.
func NewRegistryAdd(ksApp app.App, name, uri, version string, isOverride bool) (*RegistryAdd, error) {
	ra := &RegistryAdd{
		app:         ksApp,
		name:        name,
		uri:         uri,
		version:     version,
		isOverride:  isOverride,
		registryAdd: registry.Add,
	}

	return ra, nil
}

// Run lists namespaces.
func (ra *RegistryAdd) Run() error {
	uri, protocol := ra.protocol()
	_, err := ra.registryAdd(ra.app, ra.name, protocol, uri, ra.version, ra.isOverride)
	return err
}

func (ra *RegistryAdd) protocol() (string, string) {
	if strings.HasPrefix(ra.uri, "file://") {
		return ra.uri, registry.ProtocolFilesystem
	}

	if strings.HasPrefix(ra.uri, "/") {
		u := url.URL{
			Scheme: "file",
			Path:   ra.uri,
		}

		return u.String(), registry.ProtocolFilesystem
	}

	return ra.uri, registry.ProtocolGitHub
}
