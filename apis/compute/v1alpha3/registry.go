package v1alpha3

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RegistryProperties - The properties of the container registry.
type RegistryProperties struct {
	// AdminUserEnabled - The value that indicates whether the admin user is enabled.
	AdminUserEnabled bool `json:"adminUserEnabled,omitempty"`
}

// An RegistrySpec defines the desired state of a Registry.
type RegistrySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	// ResourceGroupName is the name of the resource group that the cluster will
	// be created in
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to a ResourceGroup to retrieve its
	// name
	ResourceGroupNameRef *xpv1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to a ResourceGroup to
	// retrieve its name
	ResourceGroupNameSelector *xpv1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// Sku - The SKU of the container registry.
	Sku string `json:"sku,omitempty"`

	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location string `json:"location,omitempty"`

	// RegistryProperties - The properties of the container registry.
	Properties RegistryProperties `json:",inline"`
}

// An RegistryStatus represents the observed state of an Registry.
type RegistryStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Status the status of an Azure resource at the time the operation was called.
	Status string `json:"status,omitempty"`

	// StatusMessage - The detailed message for the status, including alerts and error messages.
	StatusMessage string `json:"statusMessage,omitempty"`

	// State - The provisioning state of the container registry at the time the operation was called.
	// Possible values include: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true

// Registry an object that represents a container registry.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
// +kubebuilder:subresource:status
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegistrySpec   `json:"spec"`
	Status RegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryList contains a list of Registry.
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}
