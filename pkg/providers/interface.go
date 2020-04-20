package providers

// Bootstrap is the interface for creating a bootstrap vm and running cluster provisioning
type Bootstrap interface {
	Prepare() error
	Provision() error
	Progress() error
	Finalize() error
	Events() chan interface{}
}

// ConfigSpec holds information needed to provision a K8s management cluster
type ConfigSpec struct {
	Provider        string `yaml:"Provider" json:"provider"`
	ProviderContext Bootstrap
	SSH             SSH    `yaml:"SSH" json:"ssh"`
	LogFile         string `yaml:"LogFile,omitempty" json:"logfile,omitempty"`
	Addons          Addons `yaml:"Addons,omitempty" json:"addons,omitempty"`
	Cluster         `yaml:",inline" json:",inline" mapstructure:",squash"`
}

// Cluster specifies the details about the management cluster
type Cluster struct {
	ClusterName           string `yaml:"ClusterName" json:"clustername"`
	ControlPlaneCount     int    `yaml:"ControlPlaneCount" json:"controlplanecount"`
	ControlPlaneSize      string `yaml:"ControlPlaneSize" json:"controlplanesize"`
	WorkerCount           int    `yaml:"WorkerCount" json:"workercount"`
	WorkerSize            string `yaml:"WorkerSize" json:"workersize"`
	KubernetesVersion     string `yaml:"KubernetesVersion" json:"kubernetesversion"`
	KubernetesPodCidr     string `yaml:"KubernetesPodCidr" json:"kubernetespodcidr"`
	KubernetesServiceCidr string `yaml:"KubernetesServiceCidr" json:"kubernetesservicecidr"`
}

// SSH holds ssh info
type SSH struct {
	Username      string `yaml:"Username" json:"username"`
	AuthorizedKey string `yaml:"AuthorizedKey" json:"authorizedkey"`
}

// Addons holds optional configuration values
type Addons struct {
}
