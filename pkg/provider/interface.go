package provider

import (
	"github.com/netapp/cake/pkg/config/cluster"
	"github.com/netapp/cake/pkg/config/types"
	"github.com/netapp/cake/pkg/progress"
)

// Bootstrapper is the interface for creating infrastructure to run a cake engine against
type Bootstrapper interface {
	// Client setups up any client connections to remote provider
	Client() error
	// Prepare setups up any needed infrastructure
	Prepare() error
	// Provision runs the management cluster creation steps
	Provision() error
	// Progress watches the cluster creation for progress
	Progress() error
	// Finalize saves any deliverables and removes any created bootstrap infrastructure
	Finalize() error
	// Events are status messages from the implementation
	Events() progress.Events
}

// Spec for the Provider
type Spec struct {
	cluster.K8sConfig `yaml:",inline" json:",inline" mapstructure:",squash"`
	EventStream       progress.Events  `yaml:"-" json:"-" mapstructure:"-"`
	EngineType        types.EngineType `yaml:"EngineType" json:"enginetype"`
	LogFile           string           `yaml:"LogFile" json:"logfile"`
	LogDir            string           `yaml:"LogDir" json:"logdir"`
	SSH               cluster.SSH      `yaml:"SSH" json:"ssh"`
	BootstrapperIP    string           `yaml:"-" json:"-" mapstructure:"-"`
}

// Run provider bootstrap process
func Run(b Bootstrapper) error {
	log := b.Events()
	log.Publish(&progress.StatusEvent{
		Type:  "progress",
		Msg:   "Connecting to provider",
		Level: "info",
	})
	err := b.Client()
	if err != nil {
		return err
	}
	log.Publish(&progress.StatusEvent{
		Type:  "progress",
		Msg:   "Preparing environment",
		Level: "info",
	})
	err = b.Prepare()
	if err != nil {
		return err
	}
	log.Publish(&progress.StatusEvent{
		Type:  "progress",
		Msg:   "Provisioning cluster",
		Level: "info",
	})
	err = b.Provision()
	if err != nil {
		return err
	}
	log.Publish(&progress.StatusEvent{
		Type:  "progress",
		Msg:   "Provision Progress",
		Level: "info",
	})
	err = b.Progress()
	if err != nil {
		return err
	}
	log.Publish(&progress.StatusEvent{
		Type:  "progress",
		Msg:   "Finalizing",
		Level: "info",
	})
	err = b.Finalize()
	if err != nil {
		return err
	}
	return nil
}
