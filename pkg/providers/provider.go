package provider

import (
	v2vv1alpha1 "github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	kubevirtv1 "kubevirt.io/client-go/api/v1"
	cdiv1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
)

// Provider defines the methods required by source providers for importing a VM
type Provider interface {
	Connect(*corev1.Secret) error
	Close()
	LoadVM(v2vv1alpha1.VirtualMachineImportSourceSpec) error
	PrepareResourceMapping(*v2vv1alpha1.ResourceMappingSpec, v2vv1alpha1.VirtualMachineImportSourceSpec) error
	Validate() error
	StopVM() error
	CreateVMSpec(vmImport *v2vv1alpha1.VirtualMachineImport) *kubevirtv1.VirtualMachine
	GetDataVolumeCredentials() DataVolumeCredentials
	CreateDataVolumeMap(namespace string) map[string]cdiv1.DataVolume
	UpdateVM(vmSpec *kubevirtv1.VirtualMachine, dvs map[string]cdiv1.DataVolume)
}

// DataVolumeCredentials defines the credentials required for creating a data volume
type DataVolumeCredentials struct {
	URL           string
	CACertificate string
	KeyAccess     string
	KeySecret     string
	ConfigMapName string
	SecretName    string
}
