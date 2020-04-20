package vsphere

import (
	"github.com/vmware/govmomi/object"
)

// Resource holds information about the vSphere environment being registered
type Resource struct {
	Infrastructure
	SessionManager SessionManager
}

// Infrastructure stores information about the underlying vSphere infrastructure
type Infrastructure struct {
	Datacenter   *object.Datacenter
	Datastore    *object.Datastore
	Folder       *object.Folder
	ResourcePool *object.ResourcePool
	Network      object.NetworkReference
}

// ConfigSpec holds information needed to provision a K8s management cluster with the vsphere provider
type ConfigSpec struct {
	ProviderVsphere ProviderVsphere `yaml:"ProviderVsphere,omitempty" json:"providervsphere,omitempty"`
}

// ProviderVsphere is vsphere specifc data
type ProviderVsphere struct {
	URL               string  `yaml:"URL" json:"url"`
	Username          string  `yaml:"Username" json:"username"`
	Password          string  `yaml:"Password" json:"password"`
	Datacenter        string  `yaml:"Datacenter" json:"datacenter"`
	ResourcePool      string  `yaml:"ResourcePool" json:"resourcepool"`
	Datastore         string  `yaml:"Datastore" json:"datastore"`
	ManagementNetwork string  `yaml:"ManagementNetwork" json:"managementnetwork"`
	StorageNetwork    string  `yaml:"StorageNetwork" json:"storagenetwork"`
	OVA               OVASpec `yaml:"OVA" json:"ova"`
}

// OVASpec sets OVA information used for virtual machine templates
type OVASpec struct {
	NodeTemplate         string `yaml:"NodeTemplate" json:"nodetemplate"`
	LoadbalancerTemplate string `yaml:"LoadbalancerTemplate" json:"loadbalancertemplate"`
}
