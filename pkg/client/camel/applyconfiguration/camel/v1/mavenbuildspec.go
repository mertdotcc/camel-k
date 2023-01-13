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
	corev1 "github.com/apache/camel-k/pkg/client/camel/applyconfiguration/core/v1"
)

// MavenBuildSpecApplyConfiguration represents an declarative configuration of the MavenBuildSpec type for use
// with apply.
type MavenBuildSpecApplyConfiguration struct {
	MavenSpecApplyConfiguration `json:",inline"`
	Repositories                []RepositoryApplyConfiguration `json:"repositories,omitempty"`
	Servers                     []ServerApplyConfiguration     `json:"servers,omitempty"`
}

// MavenBuildSpecApplyConfiguration constructs an declarative configuration of the MavenBuildSpec type for use with
// apply.
func MavenBuildSpec() *MavenBuildSpecApplyConfiguration {
	return &MavenBuildSpecApplyConfiguration{}
}

// WithLocalRepository sets the LocalRepository field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LocalRepository field is set to the value of the last call.
func (b *MavenBuildSpecApplyConfiguration) WithLocalRepository(value string) *MavenBuildSpecApplyConfiguration {
	b.LocalRepository = &value
	return b
}

// WithProperties puts the entries into the Properties field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Properties field,
// overwriting an existing map entries in Properties field with the same key.
func (b *MavenBuildSpecApplyConfiguration) WithProperties(entries map[string]string) *MavenBuildSpecApplyConfiguration {
	if b.Properties == nil && len(entries) > 0 {
		b.Properties = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Properties[k] = v
	}
	return b
}

// WithSettings sets the Settings field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Settings field is set to the value of the last call.
func (b *MavenBuildSpecApplyConfiguration) WithSettings(value *ValueSourceApplyConfiguration) *MavenBuildSpecApplyConfiguration {
	b.Settings = value
	return b
}

// WithSettingsSecurity sets the SettingsSecurity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SettingsSecurity field is set to the value of the last call.
func (b *MavenBuildSpecApplyConfiguration) WithSettingsSecurity(value *ValueSourceApplyConfiguration) *MavenBuildSpecApplyConfiguration {
	b.SettingsSecurity = value
	return b
}

// WithCASecrets adds the given value to the CASecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CASecrets field.
func (b *MavenBuildSpecApplyConfiguration) WithCASecrets(values ...*corev1.SecretKeySelectorApplyConfiguration) *MavenBuildSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCASecrets")
		}
		b.CASecrets = append(b.CASecrets, *values[i])
	}
	return b
}

// WithExtension adds the given value to the Extension field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Extension field.
func (b *MavenBuildSpecApplyConfiguration) WithExtension(values ...*MavenArtifactApplyConfiguration) *MavenBuildSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExtension")
		}
		b.Extension = append(b.Extension, *values[i])
	}
	return b
}

// WithCLIOptions adds the given value to the CLIOptions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CLIOptions field.
func (b *MavenBuildSpecApplyConfiguration) WithCLIOptions(values ...string) *MavenBuildSpecApplyConfiguration {
	for i := range values {
		b.CLIOptions = append(b.CLIOptions, values[i])
	}
	return b
}

// WithRepositories adds the given value to the Repositories field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Repositories field.
func (b *MavenBuildSpecApplyConfiguration) WithRepositories(values ...*RepositoryApplyConfiguration) *MavenBuildSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRepositories")
		}
		b.Repositories = append(b.Repositories, *values[i])
	}
	return b
}

// WithServers adds the given value to the Servers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Servers field.
func (b *MavenBuildSpecApplyConfiguration) WithServers(values ...*ServerApplyConfiguration) *MavenBuildSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithServers")
		}
		b.Servers = append(b.Servers, *values[i])
	}
	return b
}
