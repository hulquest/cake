package bootstrap

import (
	"path"
	"strings"
)

// Prepare the environment for bootstrapping
func (c *Client) Prepare() error {

	templateFolder, err := c.CreateVMFolder("cake/templates")
	if err != nil {
		return err
	}

	workloadFolder, err := c.CreateVMFolder("cake/workload")
	if err != nil {
		return err
	}

	mgmtFolder, err := c.CreateVMFolder("cake/mgmt")
	if err != nil {
		return err
	}

	bootstrapFolder, err := c.CreateVMFolder("cake/bootstrap")
	if err != nil {
		return err
	}
	c.createdResources = append(
		c.createdResources,
		templateFolder["cake"],
		templateFolder["templates"],
		workloadFolder["workload"],
		mgmtFolder["mgmt"],
		bootstrapFolder["bootstrap"],
	)

	c.Folder = bootstrapFolder["bootstrap"]
	//templateName := "ubuntu-1804-kube-v1.17.3-TEST"
	templateName := strings.TrimSuffix(path.Base(c.Config.OptionalConfiguration.OVA.NodeTemplate), ".ova") + "-TEST"
	template, err := c.DeployOVATemplate(templateName, c.Config.OptionalConfiguration.OVA.NodeTemplate)
	if err != nil {
		return err
	}
	c.createdResources = append(c.createdResources, template)

	publicKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDW7BP54hSp3TrQjQq7O+oprZdXH8zbKBww/YJyCD9ksM/Y3BiFaCDwzN/vcRSslkn0kJDUq7TxmKp9bEZLTXqAiRe7GflNGoiAUuNY9EWnxt305HIkBs+OEdV6KDtnlm9sRAADflzbDi6YiMjbwNcfoRoxTgpo6BNlzv9Y3prDXiwEjxvosK+4WWIVTTEh33nNvQ5iQhPqBNgURmjQx9EDXFIRdZzA8OykPNLIqFdzmxGZWWxFbW/n6nEl/96b6w7Gx0YgzTSLs+6WAQl8SMP9l22L6puitpjihRw9cWRJ9r6x1eLqgc5Sv7gDKOMXghbmS6hy+AtrxCPPJgq7Mguc5bPAqTZlYMy98dxpHVqtAnBso/9aLOzAXX6At/0QUIwMP693B11NTGniIMtBxnD/yWvGoxTXNmXcTvj13cTzSv9czaGSJ+MTRIugtgyouZADfs8v59NV9KoaEq8umy6WEhmtw5wkjzvC5KK4N2bsM1N+8lSIKxYWxWZFsdYBP8ep442Z/2T5R8y8c5cp7tQqqapDt8JPJ0OPq3sn30BO3X8MgvmoB39j4Cqok1y9VuouPH4RalRLMR7KrASdlFengjt0vWBUoNaEuxRdJR2eOM6SpZh6YGqLdQH1MLaBOzDTH2tTLyTXCOSJpve6ZHOPbjS2BF34a1Kj52NTFtiYTw== jacob.weinstock@netapp.com"
	bootstrapVM, err := c.CloneTemplate(template, "bootstrap-vm", "", publicKey, "jacob")
	if err != nil {
		return err
	}
	c.createdResources = append(c.createdResources, bootstrapVM)

	return err
}
