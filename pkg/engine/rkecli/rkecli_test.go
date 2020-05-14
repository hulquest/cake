package rkecli

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestRKEconfig(t *testing.T) {
	t.Skip("not ready for primetime yet")
	var y map[string]interface{}
	err := yaml.Unmarshal([]byte(rawClusterYML), &y)
	if err != nil {
		t.Errorf("error unmarshaling RKE cluster config file: %s", err)
	}
	fmt.Printf("before: %v\n", y)
	y["authentication"] = map[string]interface{}{
		"sans":     []string{"172.60.0.40", "my.rancher.org"},
		"strategy": "x509",
		"webhook":  nil,
	}

	fmt.Printf("after: %v\n", y)

	_, err = yaml.Marshal(y)
	if err != nil {
		t.Errorf("error marshaling RKE cluster config file: %s", err)
	}
	//fmt.Println(string(clusterYML))
}
