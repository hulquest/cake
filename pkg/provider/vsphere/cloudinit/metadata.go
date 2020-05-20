package cloudinit

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"text/template"

	"github.com/vmware/govmomi/vim25/types"
)

// MetadataValues for cloudinit
type MetadataValues struct {
	Hostname string
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
		return nil, fmt.Errorf("unable to set cloud init metadata, %v", err)
	}

	return cloudinitMetaDataConfig, nil
}

const metadataTemplate = `
instance-id: "{{ .Hostname }}"
local-hostname: "{{ .Hostname }}"
`
