package capv

import (
	"github.com/netapp/cake/pkg/config/cluster"
	"github.com/netapp/cake/pkg/config/vsphere"
	"github.com/netapp/cake/pkg/engine"
	"github.com/netapp/cake/pkg/util/cmd"
	"os"
)

// MgmtCluster spec for CAPV
type MgmtCluster struct {
	engine.MgmtCluster      `yaml:",inline" json:",inline" mapstructure:",squash"`
	vsphere.ProviderVsphere `yaml:",inline" json:",inline" mapstructure:",squash"`
	cluster.CAPIConfig      `yaml:",inline" json:",inline" mapstructure:",squash"`
}

// Spec returns the Spec
func (m MgmtCluster) Spec() engine.MgmtCluster {
	return m.MgmtCluster
}

// NewMgmtClusterCAPV returns a new capv cluster spec
func NewMgmtClusterCAPV() *MgmtCluster {
	mc := new(MgmtCluster)
	if mc.LogFile != "" {
		cmd.FileLogLocation = mc.LogFile
		os.Truncate(mc.LogFile, 0)
	}
	return mc
}
