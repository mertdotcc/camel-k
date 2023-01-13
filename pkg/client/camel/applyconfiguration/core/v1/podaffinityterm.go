/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/apache/camel-k/pkg/client/camel/applyconfiguration/meta/v1"
)

// PodAffinityTermApplyConfiguration represents an declarative configuration of the PodAffinityTerm type for use
// with apply.
type PodAffinityTermApplyConfiguration struct {
	LabelSelector     *v1.LabelSelectorApplyConfiguration `json:"labelSelector,omitempty"`
	Namespaces        []string                            `json:"namespaces,omitempty"`
	TopologyKey       *string                             `json:"topologyKey,omitempty"`
	NamespaceSelector *v1.LabelSelectorApplyConfiguration `json:"namespaceSelector,omitempty"`
}

// PodAffinityTermApplyConfiguration constructs an declarative configuration of the PodAffinityTerm type for use with
// apply.
func PodAffinityTerm() *PodAffinityTermApplyConfiguration {
	return &PodAffinityTermApplyConfiguration{}
}

// WithLabelSelector sets the LabelSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelSelector field is set to the value of the last call.
func (b *PodAffinityTermApplyConfiguration) WithLabelSelector(value *v1.LabelSelectorApplyConfiguration) *PodAffinityTermApplyConfiguration {
	b.LabelSelector = value
	return b
}

// WithNamespaces adds the given value to the Namespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Namespaces field.
func (b *PodAffinityTermApplyConfiguration) WithNamespaces(values ...string) *PodAffinityTermApplyConfiguration {
	for i := range values {
		b.Namespaces = append(b.Namespaces, values[i])
	}
	return b
}

// WithTopologyKey sets the TopologyKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TopologyKey field is set to the value of the last call.
func (b *PodAffinityTermApplyConfiguration) WithTopologyKey(value string) *PodAffinityTermApplyConfiguration {
	b.TopologyKey = &value
	return b
}

// WithNamespaceSelector sets the NamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamespaceSelector field is set to the value of the last call.
func (b *PodAffinityTermApplyConfiguration) WithNamespaceSelector(value *v1.LabelSelectorApplyConfiguration) *PodAffinityTermApplyConfiguration {
	b.NamespaceSelector = value
	return b
}
