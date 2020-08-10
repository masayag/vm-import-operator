package controller

import "github.com/kubevirt/vm-import-operator/pkg/config"

const (
	// OsConfigMapNamespaceKey defines the configuration key for the OS mapping config map namespace
	OsConfigMapNamespaceKey = "osConfigMap.namespace"

	// OsConfigMapNameKey defines the configuration key for the OS mapping config map name
	OsConfigMapNameKey = "osConfigMap.name"
)

// ControllerConfig stores controller runtime configuration
type ControllerConfig struct {
	config.Config
}

// NewControllerConfigFrom creates new controller config from given Config
func NewControllerConfigFrom(config config.Config) ControllerConfig {
	return ControllerConfig{
		Config: config,
	}
}

// OsConfigMapNamespace provides namespace where the OS mapping ConfigMap resides. Empty string is returned when the namespace is not present.
func (c ControllerConfig) OsConfigMapNamespace() string {
	return c.ConfigMap.Data[OsConfigMapNamespaceKey]
}

// OsConfigMapName provides name of the the OS mapping ConfigMap. Empty string is returned when the name is not present.
func (c ControllerConfig) OsConfigMapName() string {
	return c.ConfigMap.Data[OsConfigMapNameKey]
}
