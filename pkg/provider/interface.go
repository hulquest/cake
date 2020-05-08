package provider

import (
	"github.com/netapp/cake/pkg/config/cluster"
	"github.com/netapp/cake/pkg/config/types"
	"github.com/netapp/cake/pkg/progress"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Bootstrapper is the interface for creating infrastructure to run a cake engine against
type Bootstrapper interface {
	// Init interface level logging
	Init()
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
	Events() chan string
}

// Spec for the Provider
type Spec struct {
	cluster.K8sConfig `yaml:",inline" json:",inline" mapstructure:",squash"`
	EventStream       chan string      `yaml:"-" json:"-" mapstructure:"-"`
	EngineType        types.EngineType `yaml:"EngineType" json:"enginetype"`
	LogFile           string           `yaml:"LogFile" json:"logfile"`
	SSH               cluster.SSH      `yaml:"SSH" json:"ssh"`
}

// Run provider bootstrap process
func Run(b Bootstrapper) error {
	writer := progress.NewChanWriter(b.Events())
	log = logrus.New()
	log.Out = writer
	log.SetFormatter(&progress.LogrusFormat{})
	b.Init()
	log.Infoln("get bootstrapper client")
	err := b.Client()
	if err != nil {
		return err
	}
	log.Infoln("prepare..")
	err = b.Prepare()
	if err != nil {
		return err
	}
	log.Infoln("provision...")
	err = b.Provision()
	if err != nil {
		return err
	}
	log.Infoln("progress...")
	err = b.Progress()
	if err != nil {
		return err
	}
	log.Infoln("finalize...")
	err = b.Finalize()
	if err != nil {
		return err
	}
	return nil
}
