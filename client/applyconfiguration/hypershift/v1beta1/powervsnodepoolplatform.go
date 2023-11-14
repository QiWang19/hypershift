/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// PowerVSNodePoolPlatformApplyConfiguration represents an declarative configuration of the PowerVSNodePoolPlatform type for use
// with apply.
type PowerVSNodePoolPlatformApplyConfiguration struct {
	SystemType        *string                                     `json:"systemType,omitempty"`
	ProcessorType     *v1beta1.PowerVSNodePoolProcType            `json:"processorType,omitempty"`
	Processors        *intstr.IntOrString                         `json:"processors,omitempty"`
	MemoryGiB         *int32                                      `json:"memoryGiB,omitempty"`
	Image             *PowerVSResourceReferenceApplyConfiguration `json:"image,omitempty"`
	StorageType       *v1beta1.PowerVSNodePoolStorageType         `json:"storageType,omitempty"`
	ImageDeletePolicy *v1beta1.PowerVSNodePoolImageDeletePolicy   `json:"imageDeletePolicy,omitempty"`
}

// PowerVSNodePoolPlatformApplyConfiguration constructs an declarative configuration of the PowerVSNodePoolPlatform type for use with
// apply.
func PowerVSNodePoolPlatform() *PowerVSNodePoolPlatformApplyConfiguration {
	return &PowerVSNodePoolPlatformApplyConfiguration{}
}

// WithSystemType sets the SystemType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SystemType field is set to the value of the last call.
func (b *PowerVSNodePoolPlatformApplyConfiguration) WithSystemType(value string) *PowerVSNodePoolPlatformApplyConfiguration {
	b.SystemType = &value
	return b
}

// WithProcessorType sets the ProcessorType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProcessorType field is set to the value of the last call.
func (b *PowerVSNodePoolPlatformApplyConfiguration) WithProcessorType(value v1beta1.PowerVSNodePoolProcType) *PowerVSNodePoolPlatformApplyConfiguration {
	b.ProcessorType = &value
	return b
}

// WithProcessors sets the Processors field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Processors field is set to the value of the last call.
func (b *PowerVSNodePoolPlatformApplyConfiguration) WithProcessors(value intstr.IntOrString) *PowerVSNodePoolPlatformApplyConfiguration {
	b.Processors = &value
	return b
}

// WithMemoryGiB sets the MemoryGiB field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MemoryGiB field is set to the value of the last call.
func (b *PowerVSNodePoolPlatformApplyConfiguration) WithMemoryGiB(value int32) *PowerVSNodePoolPlatformApplyConfiguration {
	b.MemoryGiB = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *PowerVSNodePoolPlatformApplyConfiguration) WithImage(value *PowerVSResourceReferenceApplyConfiguration) *PowerVSNodePoolPlatformApplyConfiguration {
	b.Image = value
	return b
}

// WithStorageType sets the StorageType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageType field is set to the value of the last call.
func (b *PowerVSNodePoolPlatformApplyConfiguration) WithStorageType(value v1beta1.PowerVSNodePoolStorageType) *PowerVSNodePoolPlatformApplyConfiguration {
	b.StorageType = &value
	return b
}

// WithImageDeletePolicy sets the ImageDeletePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageDeletePolicy field is set to the value of the last call.
func (b *PowerVSNodePoolPlatformApplyConfiguration) WithImageDeletePolicy(value v1beta1.PowerVSNodePoolImageDeletePolicy) *PowerVSNodePoolPlatformApplyConfiguration {
	b.ImageDeletePolicy = &value
	return b
}
