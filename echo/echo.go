package main

import (
	"flag"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/golang/glog"
	"time"
)

var port = flag.Int("port", 55555, "default port of echo server")

type EchoMessage struct {
	Message   string `json:"message,omitempty"`
	LatencyMS int    `json:"latency, omitempty"`
}

func main() {
	flag.Parse()
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			msg := new(EchoMessage)
			err := decoder.Decode(msg)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			glog.Infof(`get %v`, msg.Message)
			if msg.LatencyMS > 0 {
				d := time.Duration(msg.LatencyMS) * time.Millisecond
				glog.Infof(`sleep for: %v`, d)
				time.Sleep(d)
			}

			response, err := json.Marshal(msg)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		glog.Infof("Echo shutting down: %v", err)
	}
}
