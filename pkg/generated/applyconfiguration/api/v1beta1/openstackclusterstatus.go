/*
Copyright 2024 The Kubernetes Authors.

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
	errors "sigs.k8s.io/cluster-api-provider-openstack/pkg/utils/errors"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// OpenStackClusterStatusApplyConfiguration represents a declarative configuration of the OpenStackClusterStatus type for use
// with apply.
type OpenStackClusterStatusApplyConfiguration struct {
	Ready                     *bool                                       `json:"ready,omitempty"`
	Network                   *NetworkStatusWithSubnetsApplyConfiguration `json:"network,omitempty"`
	ExternalNetwork           *NetworkStatusApplyConfiguration            `json:"externalNetwork,omitempty"`
	Router                    *RouterApplyConfiguration                   `json:"router,omitempty"`
	APIServerLoadBalancer     *LoadBalancerApplyConfiguration             `json:"apiServerLoadBalancer,omitempty"`
	FailureDomains            *apiv1beta1.FailureDomains                  `json:"failureDomains,omitempty"`
	ControlPlaneSecurityGroup *SecurityGroupStatusApplyConfiguration      `json:"controlPlaneSecurityGroup,omitempty"`
	WorkerSecurityGroup       *SecurityGroupStatusApplyConfiguration      `json:"workerSecurityGroup,omitempty"`
	BastionSecurityGroup      *SecurityGroupStatusApplyConfiguration      `json:"bastionSecurityGroup,omitempty"`
	Bastion                   *BastionStatusApplyConfiguration            `json:"bastion,omitempty"`
	FailureReason             *errors.DeprecatedCAPIClusterStatusError    `json:"failureReason,omitempty"`
	FailureMessage            *string                                     `json:"failureMessage,omitempty"`
}

// OpenStackClusterStatusApplyConfiguration constructs a declarative configuration of the OpenStackClusterStatus type for use with
// apply.
func OpenStackClusterStatus() *OpenStackClusterStatusApplyConfiguration {
	return &OpenStackClusterStatusApplyConfiguration{}
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithReady(value bool) *OpenStackClusterStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithNetwork sets the Network field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Network field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithNetwork(value *NetworkStatusWithSubnetsApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.Network = value
	return b
}

// WithExternalNetwork sets the ExternalNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalNetwork field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithExternalNetwork(value *NetworkStatusApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.ExternalNetwork = value
	return b
}

// WithRouter sets the Router field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Router field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithRouter(value *RouterApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.Router = value
	return b
}

// WithAPIServerLoadBalancer sets the APIServerLoadBalancer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerLoadBalancer field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithAPIServerLoadBalancer(value *LoadBalancerApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.APIServerLoadBalancer = value
	return b
}

// WithFailureDomains sets the FailureDomains field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureDomains field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithFailureDomains(value apiv1beta1.FailureDomains) *OpenStackClusterStatusApplyConfiguration {
	b.FailureDomains = &value
	return b
}

// WithControlPlaneSecurityGroup sets the ControlPlaneSecurityGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControlPlaneSecurityGroup field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithControlPlaneSecurityGroup(value *SecurityGroupStatusApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.ControlPlaneSecurityGroup = value
	return b
}

// WithWorkerSecurityGroup sets the WorkerSecurityGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WorkerSecurityGroup field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithWorkerSecurityGroup(value *SecurityGroupStatusApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.WorkerSecurityGroup = value
	return b
}

// WithBastionSecurityGroup sets the BastionSecurityGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BastionSecurityGroup field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithBastionSecurityGroup(value *SecurityGroupStatusApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.BastionSecurityGroup = value
	return b
}

// WithBastion sets the Bastion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Bastion field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithBastion(value *BastionStatusApplyConfiguration) *OpenStackClusterStatusApplyConfiguration {
	b.Bastion = value
	return b
}

// WithFailureReason sets the FailureReason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureReason field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithFailureReason(value errors.DeprecatedCAPIClusterStatusError) *OpenStackClusterStatusApplyConfiguration {
	b.FailureReason = &value
	return b
}

// WithFailureMessage sets the FailureMessage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureMessage field is set to the value of the last call.
func (b *OpenStackClusterStatusApplyConfiguration) WithFailureMessage(value string) *OpenStackClusterStatusApplyConfiguration {
	b.FailureMessage = &value
	return b
}
