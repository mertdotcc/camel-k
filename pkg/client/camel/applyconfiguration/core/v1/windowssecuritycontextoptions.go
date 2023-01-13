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

// WindowsSecurityContextOptionsApplyConfiguration represents an declarative configuration of the WindowsSecurityContextOptions type for use
// with apply.
type WindowsSecurityContextOptionsApplyConfiguration struct {
	GMSACredentialSpecName *string `json:"gmsaCredentialSpecName,omitempty"`
	GMSACredentialSpec     *string `json:"gmsaCredentialSpec,omitempty"`
	RunAsUserName          *string `json:"runAsUserName,omitempty"`
	HostProcess            *bool   `json:"hostProcess,omitempty"`
}

// WindowsSecurityContextOptionsApplyConfiguration constructs an declarative configuration of the WindowsSecurityContextOptions type for use with
// apply.
func WindowsSecurityContextOptions() *WindowsSecurityContextOptionsApplyConfiguration {
	return &WindowsSecurityContextOptionsApplyConfiguration{}
}

// WithGMSACredentialSpecName sets the GMSACredentialSpecName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GMSACredentialSpecName field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithGMSACredentialSpecName(value string) *WindowsSecurityContextOptionsApplyConfiguration {
	b.GMSACredentialSpecName = &value
	return b
}

// WithGMSACredentialSpec sets the GMSACredentialSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GMSACredentialSpec field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithGMSACredentialSpec(value string) *WindowsSecurityContextOptionsApplyConfiguration {
	b.GMSACredentialSpec = &value
	return b
}

// WithRunAsUserName sets the RunAsUserName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunAsUserName field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithRunAsUserName(value string) *WindowsSecurityContextOptionsApplyConfiguration {
	b.RunAsUserName = &value
	return b
}

// WithHostProcess sets the HostProcess field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostProcess field is set to the value of the last call.
func (b *WindowsSecurityContextOptionsApplyConfiguration) WithHostProcess(value bool) *WindowsSecurityContextOptionsApplyConfiguration {
	b.HostProcess = &value
	return b
}
