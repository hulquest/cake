package progress

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	tmpfile, err := ioutil.TempFile("", "sample")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	go Serve(tmpfile.Name())
}

func shutdown() {

}

func TestDownloadTxtFile(t *testing.T) {
	err := DownloadTxtFile("http://172.60.0.68:8081/logs", "./cake_test.log")
	if err != nil {
		t.Fatal(err)
	}
}
