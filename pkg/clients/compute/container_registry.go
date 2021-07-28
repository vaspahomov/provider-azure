package compute

import (
	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"

	"github.com/crossplane/provider-azure/apis/compute/v1alpha3"
	azure "github.com/crossplane/provider-azure/pkg/clients"
)

// NewRegistry returns an Azure Registry object from a Registry spec
func NewRegistry(r v1alpha3.RegistrySpec) containerregistry.Registry {
	return containerregistry.Registry{
		Sku: &containerregistry.Sku{
			Name: containerregistry.SkuName(r.Sku),
		},
		RegistryProperties: &containerregistry.RegistryProperties{
			AdminUserEnabled: azure.ToBoolPtr(r.Properties.AdminUserEnabled),
		},
		Location: azure.ToStringPtr(r.Location),
	}
}

// RegistryUpToDate determines if a Registry is up to date
func RegistryUpToDate(r *v1alpha3.Registry, az *containerregistry.Registry) bool {
	if r.Spec.Sku != "" && r.Spec.Sku != string(az.Sku.Name) {
		return false
	}
	if r.Spec.Properties.AdminUserEnabled != azure.ToBool(az.AdminUserEnabled) {
		return false
	}
	return true
}

// RegistryInitialized determines if a Registry has been initialized
func RegistryInitialized(r *v1alpha3.Registry) bool {
	if r.Status.State == "Succeeded" || r.Status.State == "Failed" || r.Status.State == "Canceled" {
		return true
	}
	return false
}

// UpdateRegistry updates the status related to the external
// Azure Registry in the RegistryStatus
func UpdateRegistry(r *v1alpha3.Registry, az *containerregistry.Registry) {
	r.Status.State = string(az.ProvisioningState)
	if az.Status != nil {
		r.Status.Status = azure.ToString(az.Status.DisplayStatus)
		r.Status.StatusMessage = azure.ToString(az.Status.Message)
	} else {
		r.Status.Status = ""
		r.Status.StatusMessage = ""
	}
}
