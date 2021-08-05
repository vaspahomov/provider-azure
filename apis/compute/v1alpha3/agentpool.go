/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha3

import (
	"context"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networkv1alpha3 "github.com/crossplane/provider-azure/apis/network/v1alpha3"
	"github.com/crossplane/provider-azure/apis/v1alpha3"
)

// AgentPoolParameters define the desired state of an Azure Kubernetes Engine
// cluster.
type AgentPoolParameters struct {
	// ResourceGroupName is the name of the resource group that the cluster will
	// be created in
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to a ResourceGroup to retrieve its
	// name
	ResourceGroupNameRef *xpv1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to a ResourceGroup to
	// retrieve its name
	ResourceGroupNameSelector *xpv1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// AKSClusterName is the name of the AKSCluster that the cluster will
	// be created in
	AKSClusterName string `json:"aksClusterName,omitempty"`

	// AKSClusterNameRef - A reference to a AKSCluster to retrieve its
	// id
	AKSClusterNameRef *xpv1.Reference `json:"aksClusterNameRef,omitempty"`

	// AKSClusterNameSelector - Select a reference to a AKSCluster to
	// retrieve its ids
	AKSClusterNameSelector *xpv1.Selector `json:"aksClusterNameSelector,omitempty"`

	// VnetSubnetID is the subnet to which the cluster will be deployed.
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty"`

	// ResourceGroupNameRef - A reference to a Subnet to retrieve its ID
	VnetSubnetIDRef *xpv1.Reference `json:"vnetSubnetIDRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to a Subnet to retrieve
	// its ID
	VnetSubnetIDSelector *xpv1.Selector `json:"vnetSubnetIDSelector,omitempty"`

	// NodeCount is the number of nodes that the cluster will initially be
	// created with.  This can be scaled over time and defaults to 1.
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Minimum=0
	// +optional
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// MaxNodeCount - Maximum number of nodes for auto-scaling
	// +optional
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	// MinNodeCount - Minimum number of nodes for auto-scaling
	// +optional
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// EnableAutoScaling - Whether to enable auto-scaler
	EnableAutoScaling bool `json:"enableAutoScaling,omitempty"`

	// Mode - AgentPoolMode represents mode of an agent pool.
	// Possible values include: 'System', 'User'.
	// Default: 'User'
	// +optional
	Mode string `json:"mode,omitempty"`

	// AvailabilityZones - Availability zones for nodes.
	// Must use VirtualMachineScaleSets AgentPoolType.
	// +optional
	// +immutable
	AvailabilityZones []string `json:"availabilityZones,omitempty"`

	// NodeVMSize is the name of the worker node VM size, e.g., Standard_B2s,
	// Standard_F2s_v2, etc.
	// +required
	NodeVMSize string `json:"nodeVMSize"`

	// NodeTaints - Taints added to new nodes during node pool create and scale.
	// For example, key=value:NoSchedule.
	// +optional
	// +immutable
	NodeTaints []string `json:"nodeTaints,omitempty"`
}

// An AgentPoolSpec defines the desired state of a AKSCluster.
type AgentPoolSpec struct {
	xpv1.ResourceSpec   `json:",inline"`
	AgentPoolParameters `json:",inline"`
}

// An AgentPoolStatus represents the observed state of an AKSCluster.
type AgentPoolStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// State is the current state of the cluster.
	State string `json:"state,omitempty"`

	// NodesCount - Number of agents (VMs) to host docker containers.
	// Allowed values must be in the range of 0 to 100 (inclusive).
	// The default value is 1.
	NodesCount int `json:"nodesCount,omitempty"`

	// AvailabilityZones - Availability zones for nodes.
	// Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones []string `json:"availabilityZones,omitempty"`

	// ProviderID is the external ID to identify this resource in the cloud
	// provider.
	ProviderID string `json:"providerID,omitempty"`
}

// +kubebuilder:object:root=true

// An AgentPool is a managed resource that represents an Azure AgentPool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="COUNT",type="integer",JSONPath=".status.nodesCount"
// +kubebuilder:printcolumn:name="AZ",type="string",JSONPath=".status.availabilityZones"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
// +kubebuilder:subresource:status
type AgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AgentPoolSpec   `json:"spec"`
	Status AgentPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentPoolList contains a list of AgentPool.
type AgentPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentPool `json:"items"`
}

// ResolveReferences of this AgentPool.
func (mg *AgentPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.resourceGroupName
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ResourceGroupName,
		Reference:    mg.Spec.ResourceGroupNameRef,
		Selector:     mg.Spec.ResourceGroupNameSelector,
		To:           reference.To{Managed: &v1alpha3.ResourceGroup{}, List: &v1alpha3.ResourceGroupList{}},
		Extract:      reference.ExternalName(),
		Namespace:    mg.Namespace,
	})
	if err != nil {
		return errors.Wrap(err, "spec.resourceGroupName")
	}
	mg.Spec.ResourceGroupName = rsp.ResolvedValue
	mg.Spec.ResourceGroupNameRef = rsp.ResolvedReference

	// Resolve spec.vnetSubnetID
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.VnetSubnetID,
		Reference:    mg.Spec.VnetSubnetIDRef,
		Selector:     mg.Spec.VnetSubnetIDSelector,
		To:           reference.To{Managed: &networkv1alpha3.Subnet{}, List: &networkv1alpha3.SubnetList{}},
		Extract:      networkv1alpha3.SubnetID(),
		Namespace:    mg.Namespace,
	})
	if err != nil {
		return errors.Wrap(err, "spec.vnetSubnetID")
	}
	mg.Spec.VnetSubnetID = rsp.ResolvedValue
	mg.Spec.VnetSubnetIDRef = rsp.ResolvedReference

	// Resolve spec.aksClusterName
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.AKSClusterName,
		Reference:    mg.Spec.AKSClusterNameRef,
		Selector:     mg.Spec.AKSClusterNameSelector,
		To:           reference.To{Managed: &AKSCluster{}, List: &AKSClusterList{}},
		Extract:      AKSClusterName(),
		Namespace:    mg.Namespace,
	})
	if err != nil {
		return errors.Wrap(err, "spec.aksClusterName")
	}
	mg.Spec.AKSClusterName = rsp.ResolvedValue
	mg.Spec.AKSClusterNameRef = rsp.ResolvedReference

	return nil
}
