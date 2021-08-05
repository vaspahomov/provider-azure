/*
Copyright 2019 The Crossplane Authors.

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

	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networkv1alpha3 "github.com/crossplane/provider-azure/apis/network/v1alpha3"
	"github.com/crossplane/provider-azure/apis/v1alpha3"
)

// AKSClusterName extracts Name from the supplied managed resource, which must be
// a AKSCluster.
func AKSClusterName() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		s, ok := mg.(*AKSCluster)
		if !ok {
			return ""
		}
		return s.Name
	}
}

// ResolveReferences of this AKSCluster.
func (mg *AKSCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	return nil
}

// ResolveReferences of this VirtualMachine.
func (mg *VirtualMachine) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	if mg.Spec.VirtualMachineParameters != nil && mg.Spec.VirtualMachineParameters.NetworkProfile != nil {
		netProf := mg.Spec.VirtualMachineParameters.NetworkProfile
		for _, networkInterface := range netProf.NetworkInterfaces {
			// Resolve spec.properties.networkProfile.networkInterfaces[].networkInterfaceID
			rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: networkInterface.NetworkInterfaceID,
				Reference:    networkInterface.NetworkInterfaceIDRef,
				Selector:     networkInterface.NetworkInterfaceIDSelector,
				To:           reference.To{Managed: &networkv1alpha3.NetworkInterface{}, List: &networkv1alpha3.NetworkInterfaceList{}},
				Extract:      networkv1alpha3.NetworkInterfaceID(),
				Namespace:    mg.Namespace,
			})
			if err != nil {
				return errors.Wrap(err, "spec.properties.networkProfile.networkInterfaces[].networkInterfaceID")
			}
			networkInterface.NetworkInterfaceID = rsp.ResolvedValue
			networkInterface.NetworkInterfaceIDRef = rsp.ResolvedReference
		}
	}
	return nil
}

// ResolveReferences of this Registry.
func (mg *Registry) ResolveReferences(ctx context.Context, c client.Reader) error {
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
	return nil
}
