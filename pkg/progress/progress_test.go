package progress

import (
	"testing"
)

func TestDownloadTxtFile(t *testing.T) {
	err := DownloadTxtFile("http://172.60.0.68:8081/logs", "./cake_test.log")
	if err != nil {
		t.Fatal(err)
	}
}
