package capv

import (
	"github.com/netapp/cake/pkg/progress"
)

// Events returns the channel of progress messages
func (m MgmtCluster) Events() progress.Events {
	return m.EventStream
}
