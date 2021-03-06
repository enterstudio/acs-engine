package kubernetesupgrade

import (
	"os"
	"testing"

	"github.com/Azure/acs-engine/pkg/api"
	"github.com/Azure/acs-engine/pkg/armhelpers"
	"github.com/Azure/acs-engine/pkg/i18n"
	. "github.com/Azure/acs-engine/pkg/test"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo"
	uuid "github.com/satori/go.uuid"
)

func TestUpgradeCluster(t *testing.T) {
	RunSpecsWithReporters(t, "kubernetesupgrade", "Server Suite")
}

var _ = Describe("Upgrade Kubernetes cluster tests", func() {
	AfterEach(func() {
		// delete temp template directory
		os.RemoveAll("_output")
	})

	It("Should return error message when failing to list VMs during upgrade operation", func() {
		cs := api.ContainerService{}
		ucs := api.UpgradeContainerService{}

		locale, _ := i18n.LoadTranslations()
		uc := UpgradeCluster{
			Translator: &i18n.Translator{
				Locale: locale,
			},
		}

		mockClient := armhelpers.MockACSEngineClient{}
		mockClient.FailListVirtualMachines = true
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", &cs, &ucs, "12345678")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("Error while querying ARM for resources: ListVirtualMachines failed"))

		// Clean up
		os.RemoveAll("./translations")
	})

	It("Should return error message when failing to detete VMs during upgrade operation", func() {
		cs := createContainerService("testcluster", api.KubernetesRelease1Dot5, 1, 1)

		ucs := api.UpgradeContainerService{}
		ucs.OrchestratorProfile = &api.OrchestratorProfile{}
		ucs.OrchestratorProfile.OrchestratorType = api.Kubernetes
		ucs.OrchestratorProfile.OrchestratorRelease = api.KubernetesRelease1Dot6
		ucs.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[api.KubernetesRelease1Dot6]

		uc := UpgradeCluster{}

		mockClient := armhelpers.MockACSEngineClient{}
		mockClient.FailDeleteVirtualMachine = true
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", cs, &ucs, "12345678")

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("DeleteVirtualMachine failed"))
	})

	It("Should return error message when failing to deploy template during upgrade operation", func() {
		cs := createContainerService("testcluster", api.KubernetesRelease1Dot6, 1, 1)

		ucs := api.UpgradeContainerService{}
		ucs.OrchestratorProfile = &api.OrchestratorProfile{}
		ucs.OrchestratorProfile.OrchestratorType = api.Kubernetes
		ucs.OrchestratorProfile.OrchestratorRelease = api.KubernetesRelease1Dot6
		ucs.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[api.KubernetesRelease1Dot6]

		uc := UpgradeCluster{}

		mockClient := armhelpers.MockACSEngineClient{}
		mockClient.FailDeployTemplate = true
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", cs, &ucs, "12345678")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("DeployTemplate failed"))
	})

	It("Should return error message when failing to get a virtual machine during upgrade operation", func() {
		cs := createContainerService("testcluster", api.KubernetesRelease1Dot5, 1, 6)

		ucs := api.UpgradeContainerService{}
		ucs.OrchestratorProfile = &api.OrchestratorProfile{}
		ucs.OrchestratorProfile.OrchestratorType = api.Kubernetes
		ucs.OrchestratorProfile.OrchestratorRelease = api.KubernetesRelease1Dot6
		ucs.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[api.KubernetesRelease1Dot6]

		uc := UpgradeCluster{}

		mockClient := armhelpers.MockACSEngineClient{}
		mockClient.FailGetVirtualMachine = true
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", cs, &ucs, "12345678")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("GetVirtualMachine failed"))
	})

	It("Should return error message when failing to get storage client during upgrade operation", func() {
		cs := createContainerService("testcluster", api.KubernetesRelease1Dot5, 5, 1)

		ucs := api.UpgradeContainerService{}
		ucs.OrchestratorProfile = &api.OrchestratorProfile{}
		ucs.OrchestratorProfile.OrchestratorType = api.Kubernetes
		ucs.OrchestratorProfile.OrchestratorRelease = api.KubernetesRelease1Dot6
		ucs.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[api.KubernetesRelease1Dot6]

		uc := UpgradeCluster{}

		mockClient := armhelpers.MockACSEngineClient{}
		mockClient.FailGetStorageClient = true
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", cs, &ucs, "12345678")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("GetStorageClient failed"))
	})

	It("Should return error message when failing to delete network interface during upgrade operation", func() {
		cs := createContainerService("testcluster", api.KubernetesRelease1Dot5, 3, 2)

		ucs := api.UpgradeContainerService{}
		ucs.OrchestratorProfile = &api.OrchestratorProfile{}
		ucs.OrchestratorProfile.OrchestratorType = api.Kubernetes
		ucs.OrchestratorProfile.OrchestratorRelease = api.KubernetesRelease1Dot6
		ucs.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[api.KubernetesRelease1Dot6]

		uc := UpgradeCluster{}

		mockClient := armhelpers.MockACSEngineClient{}
		mockClient.FailDeleteNetworkInterface = true
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", cs, &ucs, "12345678")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("DeleteNetworkInterface failed"))
	})

	It("Should return error message when failing on ClusterPreflightCheck operation", func() {
		cs := createContainerService("testcluster", api.KubernetesRelease1Dot5, 3, 3)

		ucs := api.UpgradeContainerService{}
		ucs.OrchestratorProfile = &api.OrchestratorProfile{}
		ucs.OrchestratorProfile.OrchestratorType = api.Kubernetes
		ucs.OrchestratorProfile.OrchestratorRelease = api.KubernetesRelease1Dot7
		ucs.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[api.KubernetesRelease1Dot7]

		uc := UpgradeCluster{}

		mockClient := armhelpers.MockACSEngineClient{}
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", cs, &ucs, "12345678")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("Upgrade to Kubernetes 1.7 is not supported from orchestrator release: 1.5"))
	})

	It("Should return error message when failing on ClusterPreflightCheck operation", func() {
		cs := createContainerService("testcluster", api.KubernetesRelease1Dot7, 3, 3)

		ucs := api.UpgradeContainerService{}
		ucs.OrchestratorProfile = &api.OrchestratorProfile{}
		ucs.OrchestratorProfile.OrchestratorType = api.Kubernetes
		ucs.OrchestratorProfile.OrchestratorRelease = api.KubernetesRelease1Dot6
		ucs.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[api.KubernetesRelease1Dot6]

		uc := UpgradeCluster{}

		mockClient := armhelpers.MockACSEngineClient{}
		uc.Client = &mockClient

		subID, _ := uuid.FromString("DEC923E3-1EF1-4745-9516-37906D56DEC4")

		err := uc.UpgradeCluster(subID, "TestRg", cs, &ucs, "12345678")
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("Upgrade to Kubernetes 1.6 is not supported from orchestrator release: 1.7"))
	})
})

func createContainerService(containerServiceName string, orchestratorRelease string, masterCount int, agentCount int) *api.ContainerService {
	cs := api.ContainerService{}
	cs.ID = uuid.NewV4().String()
	cs.Location = "eastus"
	cs.Name = containerServiceName

	cs.Properties = &api.Properties{}

	cs.Properties.MasterProfile = &api.MasterProfile{}
	cs.Properties.MasterProfile.Count = masterCount
	cs.Properties.MasterProfile.DNSPrefix = "testmaster"
	cs.Properties.MasterProfile.VMSize = "Standard_D2_v2"

	cs.Properties.AgentPoolProfiles = []*api.AgentPoolProfile{}
	agentPool := &api.AgentPoolProfile{}
	agentPool.Count = agentCount
	agentPool.Name = "agentpool1"
	agentPool.VMSize = "Standard_D2_v2"
	agentPool.OSType = "Linux"
	agentPool.AvailabilityProfile = "AvailabilitySet"
	agentPool.StorageProfile = "StorageAccount"

	cs.Properties.AgentPoolProfiles = append(cs.Properties.AgentPoolProfiles, agentPool)

	cs.Properties.LinuxProfile = &api.LinuxProfile{
		AdminUsername: "azureuser",
		SSH: struct {
			PublicKeys []api.PublicKey `json:"publicKeys"`
		}{},
	}

	cs.Properties.LinuxProfile.AdminUsername = "azureuser"
	cs.Properties.LinuxProfile.SSH.PublicKeys = append(
		cs.Properties.LinuxProfile.SSH.PublicKeys, api.PublicKey{KeyData: "test"})

	cs.Properties.ServicePrincipalProfile = &api.ServicePrincipalProfile{}
	cs.Properties.ServicePrincipalProfile.ClientID = "DEC923E3-1EF1-4745-9516-37906D56DEC4"
	cs.Properties.ServicePrincipalProfile.Secret = "DEC923E3-1EF1-4745-9516-37906D56DEC4"

	cs.Properties.OrchestratorProfile = &api.OrchestratorProfile{}
	cs.Properties.OrchestratorProfile.OrchestratorType = api.Kubernetes
	cs.Properties.OrchestratorProfile.OrchestratorRelease = orchestratorRelease
	cs.Properties.OrchestratorProfile.OrchestratorVersion = api.KubernetesReleaseToVersion[orchestratorRelease]

	return &cs
}
