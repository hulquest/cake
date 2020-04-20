package bootstrap

import (
	"github.com/netapp/cake/pkg/config/types"
	"github.com/netapp/cake/pkg/providers"
	"github.com/netapp/cake/pkg/engines/capv"
	"github.com/netapp/cake/pkg/providers/vsphere"
)

const (
	uploadPort       = "50000"
	commandPort      = "50001"
	remoteExecutable = "/tmp/cake"
	remoteConfig     = "~/.cake.yaml"
)

// Client bootstrapping
type Client struct {
	vsphere.Resource
	Config       types.ConfigSpec
	resources    map[string]interface{}
	events       chan interface{}
	EngineConfig capv.MgmtCluster
}

// NewBootstrapper creates a new Bootstrap interface
func NewBootstrapper(c *Client) providers.Bootstrap {
	bc := new(Client)
	bc = c
	bc.events = make(chan interface{})
	bc.resources = make(map[string]interface{})

	return bc
}

// Event spec
type Event struct {
	EventType string
	Event     string
}
