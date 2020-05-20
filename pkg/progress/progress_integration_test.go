// +build integration

package progress

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/go-nats"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	files := setup()
	code := m.Run()
	shutdown(files)
	os.Exit(code)
}

func setup() []*os.File {
	var files []*os.File
	var err error
	logfile, err := ioutil.TempFile("", "log")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files = append(files, logfile)
	logfile.WriteString("this is the log file")

	kubefile, err := ioutil.TempFile("", "kube")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files = append(files, kubefile)
	kubefile.WriteString("this is the kubeconfig")

	fileone, err := ioutil.TempFile("", "one*.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files = append(files, fileone)
	fileone.WriteString("this is file one")

	filetwo, err := ioutil.TempFile("", "two*.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files = append(files, filetwo)
	filetwo.WriteString("this is file two")

	deliverables := []string{
		fileone.Name(),
		filetwo.Name(),
		kubefile.Name(),
	}

	go RunServer()
	events, err := NewNatsPubSub(nats.DefaultURL, "test_cluster")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go Serve(logfile.Name(), "localhost", "8081", events, deliverables)
	return files
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func shutdown(files []*os.File) {
	for _, elem := range files {
		os.Remove(elem.Name())
	}
}

func TestDownloadTxtFile(t *testing.T) {
	err := DownloadTxtFile("http://localhost:8081/log", "./cake_test.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("cake_test.log")
	ok := fileExists("cake_test.log")
	if !ok {
		t.Fatal("file not written to disk")
	}

	r, err := http.Get("http://localhost:8081/deliverable")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	var deliverables []DeliverableInfo
	json.Unmarshal(resp, &deliverables)
	for _, elem := range deliverables {
		fmt.Printf("%+v\n", elem)
		fname := fmt.Sprintf("%s%s", filepath.Base(elem.Url), elem.FileExt)
		err := DownloadTxtFile(fmt.Sprintf("http://localhost:8081%s", elem.Url), fname)
		if err != nil {
			t.Fatal(err)
		}
		ok := fileExists(fname)
		if !ok {
			t.Fatal("file not written to disk")
		}
		os.Remove(fname)
	}
}
