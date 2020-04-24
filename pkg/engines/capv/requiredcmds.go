package capv

import (
	"os"

	"github.com/netapp/cake/pkg/cmds"
)

type requiredCmd string

const (
	kind       requiredCmd = "kind"
	clusterctl requiredCmd = "clusterctl"
	kubectl    requiredCmd = "kubectl"
	docker     requiredCmd = "docker"
	helm       requiredCmd = "helm"
	tridentctl requiredCmd = "tridentctl"
)

// RequiredCommands for capv provisioner
var RequiredCommands = cmds.ProvisionerCommands{Name: "required CAPV bootstrap commands"}

// RequiredCommands checks the PATH for required commands
func (m MgmtCluster) RequiredCommands() []string {

	if m.LogFile != "" {
		cmds.FileLogLocation = m.LogFile
		os.Truncate(m.LogFile, 0)
	}

	kd := cmds.NewCommandLine(nil, string(kind), nil, nil)
	RequiredCommands.AddCommand(kd.CommandName, kd)
	c := cmds.NewCommandLine(nil, string(clusterctl), nil, nil)
	RequiredCommands.AddCommand(c.CommandName, c)
	k := cmds.NewCommandLine(nil, string(kubectl), nil, nil)
	RequiredCommands.AddCommand(k.CommandName, k)
	d := cmds.NewCommandLine(nil, string(docker), nil, nil)
	RequiredCommands.AddCommand(d.CommandName, d)

	if m.Addons.Observability.Enable {
		h := cmds.NewCommandLine(nil, string(helm), nil, nil)
		RequiredCommands.AddCommand(h.CommandName, h)
	}

	if m.Addons.Solidfire.Enable {
		t := cmds.NewCommandLine(nil, string(tridentctl), nil, nil)
		RequiredCommands.AddCommand(t.CommandName, t)
	}

	return RequiredCommands.Exist()
}
