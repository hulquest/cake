package progress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var responseBody *Status

type Status struct {
	Complete              bool     `json:"complete"`
	CompletedSuccessfully bool     `json:"completedSuccessfully"`
	Messages              []string `json:"messages"`
}

func UpdateProgressComplete(complete bool) {
	responseBody.Complete = complete
}

func UpdateProgressCompletedSuccessfully(completedSuccessfully bool) {
	responseBody.CompletedSuccessfully = completedSuccessfully
}

func init() {
	responseBody = new(Status)
	responseBody.Messages = []string{}
}

func Serve(logfile string, kubeconfig string, port string, status Events) {
	fn := func(p *StatusEvent) {
		responseBody.Messages = append(responseBody.Messages, fmt.Sprintf("%v", p.String()))
	}
	status.Subscribe(fn)

	http.HandleFunc("/progress", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(responseBody)
	})
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		logs, _ := ioutil.ReadFile(logfile)
		fmt.Fprintf(w, string(logs))
	})
	http.HandleFunc("/kubeconfig", func(w http.ResponseWriter, r *http.Request) {
		kconfig, _ := ioutil.ReadFile(kubeconfig)
		if len(kconfig) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, string(kconfig))
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ServeDuration() {
	for x := 0; x <= 24; x++ {
		time.Sleep(1 * time.Hour)
	}
}
