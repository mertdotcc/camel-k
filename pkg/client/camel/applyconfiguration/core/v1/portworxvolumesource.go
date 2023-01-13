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

// PortworxVolumeSourceApplyConfiguration represents an declarative configuration of the PortworxVolumeSource type for use
// with apply.
type PortworxVolumeSourceApplyConfiguration struct {
	VolumeID *string `json:"volumeID,omitempty"`
	FSType   *string `json:"fsType,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
}

// PortworxVolumeSourceApplyConfiguration constructs an declarative configuration of the PortworxVolumeSource type for use with
// apply.
func PortworxVolumeSource() *PortworxVolumeSourceApplyConfiguration {
	return &PortworxVolumeSourceApplyConfiguration{}
}

// WithVolumeID sets the VolumeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeID field is set to the value of the last call.
func (b *PortworxVolumeSourceApplyConfiguration) WithVolumeID(value string) *PortworxVolumeSourceApplyConfiguration {
	b.VolumeID = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *PortworxVolumeSourceApplyConfiguration) WithFSType(value string) *PortworxVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *PortworxVolumeSourceApplyConfiguration) WithReadOnly(value bool) *PortworxVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
