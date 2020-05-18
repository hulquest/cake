package cloudinit

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"text/template"

	"github.com/vmware/govmomi/vim25/types"
)

// MetadataValues used
type MetadataValues struct {
	Hostname string
	//Networks []NetworkConfig
}

// NetworkConfig values
type NetworkConfig struct {
	MACAddress  string
	DHCP4       bool
	IPAddress   string
	Netmask     string
	Gateway     string
	NameServers []string
	DNSSearch   []string
}

// SetCloudInitMetadata sets the cloud init user data at the key
// "guestinfo.metadata" as a base64-encoded string.
func (e *Config) SetCloudInitMetadata(data []byte) error {
	*e = append(*e,
		&types.OptionValue{
			Key:   "guestinfo.metadata",
			Value: base64.StdEncoding.EncodeToString(data),
		},
		&types.OptionValue{
			Key:   "guestinfo.metadata.encoding",
			Value: "base64",
		},
	)

	return nil
}

// GetMetadata returns the metadata
func GetMetadata(metadataValues *MetadataValues) ([]byte, error) {
	textTemplate, err := template.New("f").Parse(metadataTemplate)
	if err != nil {
		return nil, fmt.Errorf("unable to parse cloud init metadata template, %v", err)
	}
	returnScript := new(bytes.Buffer)
	err = textTemplate.Execute(returnScript, metadataValues)
	if err != nil {
		return nil, fmt.Errorf("unable to template cloud init metadata, %v", err)
	}

	return returnScript.Bytes(), nil
}

// GenerateMetaData creates the meta data
func GenerateMetaData(hostname string) (Config, error) {
	// Create metadata
	metadataValues := &MetadataValues{
		Hostname: hostname,
	}

	metadata, err := GetMetadata(metadataValues)
	if err != nil {
		return nil, fmt.Errorf("unable to get cloud metadata, %v", err)
	}

	var cloudinitMetaDataConfig Config

	err = cloudinitMetaDataConfig.SetCloudInitMetadata(metadata)
	if err != nil {
		return nil, fmt.Errorf("unable to set cloud init metadata in extra config, %v", err)
	}

	return cloudinitMetaDataConfig, nil
}

const metadataTemplate = `
instance-id: "{{ .Hostname }}"
local-hostname: "{{ .Hostname }}"
`

// NOTE: Debian 9 does not support v2 of cloud-init networking configuration, needs netplan.io. Using ENI configuration.
const metadataTemplatev1 = `
instance-id: "{{ .Hostname }}"
local-hostname: "{{ .Hostname }}"
network:
  version: 1
  config:
  {{- range $index, $network := .Networks}}
    - type: physical
      name: id{{ $index }}
      mac_address: {{ $network.MACAddress }}
      subnets:
	  {{- if $network.DHCP4 }}
        - type: dhcp
      {{- end }}
	  {{- if not $network.DHCP4 }}
        - type: static
          address: {{$network.IPAddress}}
          {{- if $network.Netmask }}
          netmask: {{ $network.Netmask }}
          {{- end }}
          {{- if $network.Gateway }}
          gateway: {{$network.Gateway}}
          {{- end }}
          {{- if $network.NameServers }}
          dns_nameservers:
          {{- range $network.NameServers }}
            - {{ . }}
          {{- end }}
          {{- end }}
          {{- if $network.DNSSearch }}
          dns_search:
          {{- range $network.DNSSearch }}
            - {{ . }}
          {{- end }}
          {{- end }}
      {{- end }}
  {{- end }}
`
