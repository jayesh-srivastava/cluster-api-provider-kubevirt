/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"

// Conditions and condition Reasons for the KubevirtMachine object

const (
	// VMProvisionedCondition documents the status of the provisioning of the VM
	// generated by a KubevirtMachine.
	VMProvisionedCondition clusterv1.ConditionType = "VMProvisioned"

	// WaitingForClusterInfrastructureReason (Severity=Info) documents a KubevirtMachine waiting for the cluster
	// infrastructure to be ready before starting to create the container that provides the KubevirtMachine
	// infrastructure.
	WaitingForClusterInfrastructureReason = "WaitingForClusterInfrastructure"

	// WaitingForBootstrapDataReason (Severity=Info) documents a KubevirtMachine waiting for the bootstrap
	// script to be ready before starting to create the VM that provides the KubevirtMachine infrastructure.
	WaitingForBootstrapDataReason = "WaitingForBootstrapData"
)

const (
	// BootstrapExecSucceededCondition provides an observation of the KubevirtMachine bootstrap process.
	// 	It is set based on successful execution of bootstrap commands and on the existence of
	//	the /run/cluster-api/bootstrap-success.complete file.
	// The condition gets generated after VMProvisionedCondition is True.
	BootstrapExecSucceededCondition clusterv1.ConditionType = "BootstrapExecSucceeded"

	// BootstrappingReason documents (Severity=Info) a KubevirtMachine currently executing the bootstrap
	// script that creates the Kubernetes node on the newly provisioned machine infrastructure.
	BootstrappingReason = "Bootstrapping"

	// BootstrapFailedReason documents (Severity=Warning) a KubevirtMachine controller detecting an error while
	// bootstrapping the Kubernetes node on the machine just provisioned; those kind of errors are usually
	// transient and failed bootstrap are automatically re-tried by the controller.
	BootstrapFailedReason = "BootstrapFailed"
)

// Conditions and condition Reasons for the KubevirtCluster object

const (
	// LoadBalancerAvailableCondition documents the availability of the service that implements the cluster load balancer.
	LoadBalancerAvailableCondition clusterv1.ConditionType = "LoadBalancerAvailable"

	// LoadBalancerProvisioningFailedReason (Severity=Warning) documents a KubevirtCluster controller detecting
	// an error while provisioning the service that provides the cluster load balancer; those kind of
	// errors are usually transient and failed provisioning are automatically re-tried by the controller.
	LoadBalancerProvisioningFailedReason = "LoadBalancerProvisioningFailed"
)
