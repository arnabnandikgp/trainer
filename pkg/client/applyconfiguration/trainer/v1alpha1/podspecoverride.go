// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// PodSpecOverrideApplyConfiguration represents a declarative configuration of the PodSpecOverride type for use
// with apply.
type PodSpecOverrideApplyConfiguration struct {
	TargetJobs         []PodSpecOverrideTargetJobApplyConfiguration `json:"targetJobs,omitempty"`
	Containers         []ContainerOverrideApplyConfiguration        `json:"containers,omitempty"`
	InitContainers     []ContainerOverrideApplyConfiguration        `json:"initContainers,omitempty"`
	Volumes            []v1.VolumeApplyConfiguration                `json:"volumes,omitempty"`
	ServiceAccountName *string                                      `json:"serviceAccountName,omitempty"`
	NodeSelector       map[string]string                            `json:"nodeSelector,omitempty"`
	Tolerations        []v1.TolerationApplyConfiguration            `json:"tolerations,omitempty"`
}

// PodSpecOverrideApplyConfiguration constructs a declarative configuration of the PodSpecOverride type for use with
// apply.
func PodSpecOverride() *PodSpecOverrideApplyConfiguration {
	return &PodSpecOverrideApplyConfiguration{}
}

// WithTargetJobs adds the given value to the TargetJobs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TargetJobs field.
func (b *PodSpecOverrideApplyConfiguration) WithTargetJobs(values ...*PodSpecOverrideTargetJobApplyConfiguration) *PodSpecOverrideApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTargetJobs")
		}
		b.TargetJobs = append(b.TargetJobs, *values[i])
	}
	return b
}

// WithContainers adds the given value to the Containers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Containers field.
func (b *PodSpecOverrideApplyConfiguration) WithContainers(values ...*ContainerOverrideApplyConfiguration) *PodSpecOverrideApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithContainers")
		}
		b.Containers = append(b.Containers, *values[i])
	}
	return b
}

// WithInitContainers adds the given value to the InitContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InitContainers field.
func (b *PodSpecOverrideApplyConfiguration) WithInitContainers(values ...*ContainerOverrideApplyConfiguration) *PodSpecOverrideApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithInitContainers")
		}
		b.InitContainers = append(b.InitContainers, *values[i])
	}
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *PodSpecOverrideApplyConfiguration) WithVolumes(values ...*v1.VolumeApplyConfiguration) *PodSpecOverrideApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumes")
		}
		b.Volumes = append(b.Volumes, *values[i])
	}
	return b
}

// WithServiceAccountName sets the ServiceAccountName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountName field is set to the value of the last call.
func (b *PodSpecOverrideApplyConfiguration) WithServiceAccountName(value string) *PodSpecOverrideApplyConfiguration {
	b.ServiceAccountName = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *PodSpecOverrideApplyConfiguration) WithNodeSelector(entries map[string]string) *PodSpecOverrideApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *PodSpecOverrideApplyConfiguration) WithTolerations(values ...*v1.TolerationApplyConfiguration) *PodSpecOverrideApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}
