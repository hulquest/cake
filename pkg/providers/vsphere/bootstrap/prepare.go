package bootstrap

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Prepare the environment for bootstrapping
func (c *Client) Prepare() error {

	templateFolder, err := c.CreateVMFolder("cake/templates")
	if err != nil {
		return err
	}
	c.resources["cakeFolder"] = templateFolder["cake"]
	c.resources["templatesFolder"] = templateFolder["templates"]

	workloadFolder, err := c.CreateVMFolder("cake/workloads")
	if err != nil {
		return err
	}
	c.resources["workloadsFolder"] = workloadFolder["workloads"]

	mgmtFolder, err := c.CreateVMFolder("cake/mgmt")
	if err != nil {
		return err
	}
	c.resources["mgmtFolder"] = mgmtFolder["mgmt"]

	bootstrapFolder, err := c.CreateVMFolder("cake/bootstrap")
	if err != nil {
		return err
	}
	c.resources["bootstrapFolder"] = bootstrapFolder["bootstrap"]
	c.Folder = templateFolder["templates"]

	ovas, err := c.DeployOVATemplates(c.Config.OptionalConfiguration.OVA.NodeTemplate, c.Config.OptionalConfiguration.OVA.LoadbalancerTemplate)

	c.resources[c.Config.OptionalConfiguration.OVA.NodeTemplate] = ovas[c.Config.OptionalConfiguration.OVA.NodeTemplate]
	c.resources[c.Config.OptionalConfiguration.OVA.LoadbalancerTemplate] = ovas[c.Config.OptionalConfiguration.OVA.LoadbalancerTemplate]

	c.Folder = bootstrapFolder["bootstrap"]
	publicKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDW7BP54hSp3TrQjQq7O+oprZdXH8zbKBww/YJyCD9ksM/Y3BiFaCDwzN/vcRSslkn0kJDUq7TxmKp9bEZLTXqAiRe7GflNGoiAUuNY9EWnxt305HIkBs+OEdV6KDtnlm9sRAADflzbDi6YiMjbwNcfoRoxTgpo6BNlzv9Y3prDXiwEjxvosK+4WWIVTTEh33nNvQ5iQhPqBNgURmjQx9EDXFIRdZzA8OykPNLIqFdzmxGZWWxFbW/n6nEl/96b6w7Gx0YgzTSLs+6WAQl8SMP9l22L6puitpjihRw9cWRJ9r6x1eLqgc5Sv7gDKOMXghbmS6hy+AtrxCPPJgq7Mguc5bPAqTZlYMy98dxpHVqtAnBso/9aLOzAXX6At/0QUIwMP693B11NTGniIMtBxnD/yWvGoxTXNmXcTvj13cTzSv9czaGSJ+MTRIugtgyouZADfs8v59NV9KoaEq8umy6WEhmtw5wkjzvC5KK4N2bsM1N+8lSIKxYWxWZFsdYBP8ep442Z/2T5R8y8c5cp7tQqqapDt8JPJ0OPq3sn30BO3X8MgvmoB39j4Cqok1y9VuouPH4RalRLMR7KrASdlFengjt0vWBUoNaEuxRdJR2eOM6SpZh6YGqLdQH1MLaBOzDTH2tTLyTXCOSJpve6ZHOPbjS2BF34a1Kj52NTFtiYTw== jacob.weinstock@netapp.com"
	c.EngineConfig.LogFile = "/tmp/cake.log"
	configYaml, err := yaml.Marshal(&c.EngineConfig)
	if err != nil {
		return err
	}
	script := fmt.Sprintf(`#!/bin/bash

wget -O /usr/local/bin/clusterctl https://github.com/kubernetes-sigs/cluster-api/releases/download/v0.3.0/clusterctl-$(uname | tr '[:upper:]' '[:lower:]')-amd64
chmod +x /usr/local/bin/clusterctl

wget -O /usr/local/bin/kind https://kind.sigs.k8s.io/dl/v0.7.0/kind-$(uname)-amd64
chmod +x /usr/local/bin/kind

curl https://get.docker.com/ | bash

echo -e '%s' > %s

socat -u TCP-LISTEN:%s,fork CREATE:%s,group=root,perm=0755 & disown
socat TCP-LISTEN:%s,reuseaddr,fork EXEC:"/bin/bash" & disown


`, configYaml, remoteConfig, uploadPort, remoteExecutable, commandPort)
	bootstrapVM, err := c.CloneTemplate(ovas[c.Config.OptionalConfiguration.OVA.NodeTemplate], "bootstrap-vm", script, publicKey, "jacob")
	if err != nil {
		return err
	}
	c.resources["bootstrapVM"] = bootstrapVM

	return err
}
