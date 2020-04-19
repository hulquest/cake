package providers

// Bootstrap is the interface for creating a bootstrap vm and running cluster provisioning
type Bootstrap interface {
	Prepare() error
	Provision() error
	Progress() error
	Finalize() error
	Events() chan interface{}
}
