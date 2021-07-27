package compute

import (
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2020-03-01/containerservice"
	"github.com/Azure/go-autorest/autorest/to"

	azure "github.com/crossplane/provider-azure/pkg/clients"

	"github.com/crossplane/provider-azure/apis/compute/v1alpha3"
)

// NewAgentPool returns an Azure AgentPool object from a AgentPool spec
func NewAgentPool(c *v1alpha3.AgentPool) containerservice.AgentPool {
	nodeCount := int32(v1alpha3.DefaultNodeCount)
	if c.Spec.NodeCount != nil {
		nodeCount = *c.Spec.NodeCount
	}

	return containerservice.AgentPool{
		ManagedClusterAgentPoolProfileProperties: &containerservice.ManagedClusterAgentPoolProfileProperties{
			Count:        &nodeCount,
			VMSize:       containerservice.VMSizeTypes(c.Spec.NodeVMSize),
			VnetSubnetID: to.StringPtr(c.Spec.VnetSubnetID),
			Mode:         containerservice.User,
			Type:         containerservice.VirtualMachineScaleSets,
		},
	}
}

func newManagedClusterAgentPoolProfile(c *v1alpha3.AKSCluster) containerservice.ManagedClusterAgentPoolProfile {
	nodeCount := int32(v1alpha3.DefaultNodeCount)
	if c.Spec.NodeCount != nil {
		nodeCount = int32(*c.Spec.NodeCount)
	}
	p := containerservice.ManagedClusterAgentPoolProfile{
		Name:   to.StringPtr(AgentPoolProfileName),
		Count:  &nodeCount,
		VMSize: containerservice.VMSizeTypes(c.Spec.NodeVMSize),
		Mode:   containerservice.System,
		Type:   containerservice.VirtualMachineScaleSets,
	}
	if c.Spec.VnetSubnetID != "" {
		p.VnetSubnetID = to.StringPtr(c.Spec.VnetSubnetID)
	}
	return p
}

// AgentPoolNeedUpdate determines if a AgentPool need to be updated
func AgentPoolNeedUpdate(c *v1alpha3.AgentPool, az *containerservice.AgentPool) bool {
	if azure.ToInt(c.Spec.NodeCount) != azure.ToInt(az.Count) {
		return true
	}
	if c.Spec.NodeVMSize != string(az.VMSize) {
		return true
	}
	return false
}
