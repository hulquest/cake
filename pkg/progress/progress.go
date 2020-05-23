package progress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Constants block for endpoint names.
const (
	URIProgress    = "/progress"
	URILogs        = "/log"
	URIDeliverable = "/deliverable"
)

var responseBody *Status

// Status is a struct to model progress.
type Status struct {
	Complete              bool     `json:"complete"`
	CompletedSuccessfully bool     `json:"completedSuccessfully"`
	Messages              []string `json:"messages"`
}

// DeliverableInfo is a stuct for 
type DeliverableInfo struct {
	URL     string `json:"url"`
	FileExt string `json:"file_extension"`
}

// UpdateProgressComplete will complete progress update.
func UpdateProgressComplete(complete bool) {
	responseBody.Complete = complete
}

// UpdateProgressCompletedSuccessfully will finalize progress.
func UpdateProgressCompletedSuccessfully(completedSuccessfully bool) {
	responseBody.CompletedSuccessfully = completedSuccessfully
}

func init() {
	responseBody = new(Status)
	responseBody.Messages = []string{}
}

// Serve will start an http server to handle events as the deployment progresses.
func Serve(logfile string, ip, port string, status Events, fileDeliverables []string) {
	fullURL := url.URL{Scheme: "http", Host: ip + ":" + port, Path: ""}
	fn := func(p *StatusEvent) {
		responseBody.Messages = append(responseBody.Messages, fmt.Sprintf("%v", p.String()))
	}
	if status != nil {
		status.Subscribe(fn)
	}

	http.HandleFunc(URIProgress, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(responseBody)
	})
	http.HandleFunc(URILogs, func(w http.ResponseWriter, r *http.Request) {
		logs, _ := ioutil.ReadFile(logfile)
		fmt.Fprintf(w, string(logs))
	})
	fullURL.Path = URILogs
	status.Publish(&StatusEvent{
		Type:  "progress",
		Msg:   fmt.Sprintf("serving file: %v at %v", logfile, fullURL.String()),
		Level: "debug",
	})

	var dv []DeliverableInfo
	for _, elem := range fileDeliverables {
		var f string
		if elem != "" {
			f = elem
			base := strings.TrimSuffix(filepath.Base(elem), filepath.Ext(elem))
			uri := fmt.Sprintf("%s/%s", URIDeliverable, base)
			http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
				file, _ := ioutil.ReadFile(f)
				fmt.Fprintf(w, string(file))
			})
			dv = append(dv, DeliverableInfo{
				URL:     uri,
				FileExt: filepath.Ext(elem),
			})
			fullURL.Path = uri
			status.Publish(&StatusEvent{
				Type:  "progress",
				Msg:   fmt.Sprintf("serving file: %v at %v", elem, fullURL.String()),
				Level: "debug",
			})
		}
	}
	http.HandleFunc(URIDeliverable, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(dv)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// ServeDuration will sleep for 25 hours?
func ServeDuration() {
	for x := 0; x <= 24; x++ {
		time.Sleep(1 * time.Hour)
	}
}

// DownloadTxtFile will download a text file from a URL to the download location.
func DownloadTxtFile(url string, downloadLocation string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error with GET on: %v, err: %v", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed, %v", resp.StatusCode)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	err = ioutil.WriteFile(downloadLocation, responseData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file to disk: %v, err: %v", downloadLocation, err)
	}
	return nil
}
