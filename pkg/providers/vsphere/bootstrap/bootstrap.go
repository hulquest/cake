package bootstrap

import (
	"github.com/netapp/cake/pkg/providers"
	"github.com/netapp/cake/pkg/providers/vsphere"
)

// Client bootstrapping
type Client struct {
	vsphere.Resource
	Config interface{}
	createdResources []interface{}
	events chan interface{}
}

// NewBootstrapper creates a new Bootstrap interface
func NewBootstrapper(c *Client) providers.Bootstrap {
	bc := new(Client)
	bc = c
	bc.events = make(chan interface{})
	
	return bc
}

// Event spec
type Event struct {
	EventType string
	Event     string
}
