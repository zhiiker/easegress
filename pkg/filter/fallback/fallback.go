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

package fallback

import (
	"github.com/megaease/easegress/pkg/context"
	"github.com/megaease/easegress/pkg/object/httppipeline"
	"github.com/megaease/easegress/pkg/supervisor"
	"github.com/megaease/easegress/pkg/util/fallback"
)

const (
	// Kind is the kind of Fallback.
	Kind = "Fallback"

	resultFallback = "fallback"
)

var results = []string{resultFallback}

func init() {
	httppipeline.Register(&Fallback{})
}

type (
	// Fallback is filter Fallback.
	Fallback struct {
		super    *supervisor.Supervisor
		pipeSpec *httppipeline.FilterSpec
		spec     *Spec

		f *fallback.Fallback
	}

	// Spec describes the Fallback.
	Spec struct {
		fallback.Spec `yaml:",inline"`
	}
)

// Kind returns the kind of Fallback.
func (f *Fallback) Kind() string {
	return Kind
}

// DefaultSpec returns default spec of Fallback.
func (f *Fallback) DefaultSpec() interface{} {
	return &Spec{}
}

// Description returns the description of Fallback.
func (f *Fallback) Description() string {
	return "Fallback do the fallback."
}

// Results returns the results of Fallback.
func (f *Fallback) Results() []string {
	return results
}

// Init initializes Fallback.
func (f *Fallback) Init(pipeSpec *httppipeline.FilterSpec, super *supervisor.Supervisor) {
	f.pipeSpec, f.spec, f.super = pipeSpec, pipeSpec.FilterSpec().(*Spec), super
	f.reload()
}

// Inherit inherits previous generation of Fallback.
func (f *Fallback) Inherit(pipeSpec *httppipeline.FilterSpec,
	previousGeneration httppipeline.Filter, super *supervisor.Supervisor) {

	previousGeneration.Close()
	f.Init(pipeSpec, super)
}

func (f *Fallback) reload() {
	f.f = fallback.New(&f.spec.Spec)
}

// Handle fallabcks HTTPContext.
// It always returns fallback.
func (f *Fallback) Handle(ctx context.HTTPContext) string {
	f.f.Fallback(ctx)
	return ctx.CallNextHandler(resultFallback)
}

// Status returns Status.
func (f *Fallback) Status() interface{} {
	return nil
}

// Close closes Fallback.
func (f *Fallback) Close() {}
