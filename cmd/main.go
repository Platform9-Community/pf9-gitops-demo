package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const PORT = 8080

func main() {
	startServer(handler)
}

func startServer(handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/", handler)
	log.Printf("starting server...")
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./version.json")
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

	log.Printf("Service (version: %s): received request from %s", obj.VERSION, r.Header.Get("User-Agent"))
	host, err := os.Hostname()
	if err != nil {
		host = "unknown host"
	}
	resp := fmt.Sprintf("Hello %s from service (version: %s) running on host %s! Go to https://platform9.com/signup/ today to create a free PMK account!!", r.Header.Get("User-Agent"), obj.VERSION, host)
	_, err = w.Write([]byte(resp))
	if err != nil {
		log.Panicf("not able to write http output: %s", err)
	}
}
