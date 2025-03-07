/*
 * Copyright (c) 2017, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package function

import (
	"fmt"

	"github.com/megaease/easegress/pkg/logger"
	"github.com/megaease/easegress/pkg/object/function/spec"
	"github.com/megaease/easegress/pkg/object/function/worker"
	"github.com/megaease/easegress/pkg/supervisor"
	"github.com/megaease/easegress/pkg/v"
	"gopkg.in/yaml.v2"
)

const (
	// Category is the category of FaasController.
	Category = supervisor.CategoryBusinessController

	// Kind is the kind of FaaSController.
	Kind = "FaaSController"
)

func init() {
	supervisor.Register(&FaasController{})
}

type (
	// FaasController is Function controller.
	FaasController struct {
		super     *supervisor.Supervisor
		superSpec *supervisor.Spec
		spec      *spec.Admin

		worker *worker.Worker
	}
)

// Category returns the category of FaasController.
func (f *FaasController) Category() supervisor.ObjectCategory {
	return Category
}

// Kind returns the kind of FaasController.
func (f *FaasController) Kind() string {
	return Kind
}

// DefaultSpec returns the default spec of Function.
func (f *FaasController) DefaultSpec() interface{} {
	return &spec.Admin{
		SyncInterval: "10s",
		Provider:     spec.ProviderKnative,
		Knative: &spec.Knative{
			Namespace: "default",
			Timeout:   "2s",
		},
	}
}

func (f *FaasController) Validate() error {
	switch f.spec.Provider {
	case spec.ProviderKnative:
		//
	default:
		return fmt.Errorf("unknown faas provider: %s", f.spec.Provider)
	}

	buff, err := yaml.Marshal(f.spec.HTTPServer)
	if err != nil {
		err = fmt.Errorf("BUG: marshal %#v to yaml failed: %v",
			f.spec.HTTPServer, err)
		logger.Errorf(err.Error())
		return err
	}

	vr := v.Validate(f.spec.HTTPServer, buff)
	if !vr.Valid() {
		return fmt.Errorf("%s", vr.Error())
	}
	return nil
}

// Init initializes Function.
func (f *FaasController) Init(superSpec *supervisor.Spec, super *supervisor.Supervisor) {
	f.superSpec, f.spec, f.super = superSpec, superSpec.ObjectSpec().(*spec.Admin), super
	f.reload()
}

// Inherit inherits previous generation of Function.
func (f *FaasController) Inherit(superSpec *supervisor.Spec,
	previousGeneration supervisor.Object, super *supervisor.Supervisor) {

	previousGeneration.Close()
	f.Init(superSpec, super)
}

func (f *FaasController) reload() {
	f.worker = worker.NewWorker(f.superSpec, f.super)
}

// Status returns Status generated by Runtime.
func (f *FaasController) Status() *supervisor.Status {
	return &supervisor.Status{}
}

// Close closes Function.
func (f *FaasController) Close() {
	f.worker.Close()
}
