package main

// add fix aaa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

const (
	srvAddr = ":8080"
)

func main() {
	data, err := ioutil.ReadFile("./package.json")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	type AppVersion struct {
		VERSION string
	}
	var obj AppVersion
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	host, err := os.Hostname()
	if err != nil {
		host = "unknown host"
	}
	// Create our middleware.
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	if err != nil {
		log.Panicf("not able to write http output: %s", err)
	}
	// Create our server.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		resp := fmt.Sprintf("<html><body><h3>%s<br>version <a href='https://github.com/Platform9-Community/pf9-gitops-demo/releases/tag/%s'>%s</a> on host %s<br></h3><a href='https://platform9.com/signup/'><img src='https://platform9.com/wp-content/uploads/2021/11/platform9_open-distributed-cloud-diagram.png'></a></body></html>", r.Header.Get("User-Agent"), obj.VERSION, obj.VERSION, host)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(resp))
		log.Printf("Service (version: %s): received request from %s", obj.VERSION, r.Header.Get("User-Agent"))
	})
	mux.Handle("/metrics", promhttp.Handler())

	// Wrap our main handler, we pass empty handler ID so the middleware infers
	// the handler label from the URL.
	h := std.Handler("", mdlw, mux)

	// Serve our handler.
	go func() {
		log.Printf("server listening at %s", srvAddr)
		if err := http.ListenAndServe(srvAddr, h); err != nil {
			log.Panicf("error while serving: %s", err)
		}
	}()

	// Wait until some signal is captured.
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
	<-sigC
}
