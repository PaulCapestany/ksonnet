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
	"testing"

	"github.com/ksonnet/ksonnet/component"
	cmocks "github.com/ksonnet/ksonnet/component/mocks"
	"github.com/ksonnet/ksonnet/metadata/app"
	amocks "github.com/ksonnet/ksonnet/metadata/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParamSet(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		componentName := "deployment"
		path := "replicas"
		value := "3"

		cm := &cmocks.Manager{}

		var ns component.Component
		c := &cmocks.Component{}
		c.On("SetParam", []string{"replicas"}, 3, component.ParamOptions{}).Return(nil)

		cm.On("ResolvePath", appMock, "deployment").Return(ns, c, nil)

		a, err := NewParamSet(appMock, componentName, path, value)
		require.NoError(t, err)

		a.cm = cm

		err = a.Run()
		require.NoError(t, err)
	})
}

func TestParamSet_index(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		componentName := "deployment"
		path := "replicas"
		value := "3"

		cm := &cmocks.Manager{}

		var ns component.Component
		c := &cmocks.Component{}
		c.On("SetParam", []string{"replicas"}, 3, component.ParamOptions{Index: 1}).Return(nil)

		cm.On("ResolvePath", appMock, "deployment").Return(ns, c, nil)

		idxOpt := ParamSetWithIndex(1)

		a, err := NewParamSet(appMock, componentName, path, value, idxOpt)
		require.NoError(t, err)

		a.cm = cm

		err = a.Run()
		require.NoError(t, err)
	})
}

func TestParamSet_global(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		nsName := "/"
		path := "replicas"
		value := "3"

		cm := &cmocks.Manager{}

		ns := &cmocks.Namespace{}
		ns.On("SetParam", []string{"replicas"}, 3).Return(nil)

		cm.On("Namespace", appMock, "/").Return(ns, nil)

		gOpt := ParamSetGlobal(true)
		a, err := NewParamSet(appMock, nsName, path, value, gOpt)
		require.NoError(t, err)

		a.cm = cm

		err = a.Run()
		require.NoError(t, err)
	})
}

func TestParamSet_env(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		name := "deployment"
		path := "replicas"
		value := "3"

		envOpt := ParamSetEnv("default")
		a, err := NewParamSet(appMock, name, path, value, envOpt)
		require.NoError(t, err)

		envSetter := func(ksApp app.App, envName, name, pName, value string) error {
			assert.Equal(t, "default", envName)
			assert.Equal(t, "deployment", name)
			assert.Equal(t, "replicas", pName)
			assert.Equal(t, "3", value)
			return nil
		}
		a.setEnv = envSetter

		err = a.Run()
		require.NoError(t, err)
	})
}
