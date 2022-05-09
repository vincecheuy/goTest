package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var delay int = 0
var ptGetAccounts, err = ioutil.ReadFile("../static/getaccounts.json")

func main() {
	fmt.Println(string(ptGetAccounts))
	fmt.Println(err)
	http.DefaultServeMux.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	headers := w.Header()
	headers["Content-Type"] = []string{"application/json"}
	path := r.URL.Path
	if strings.HasPrefix(path, "/pt-mock/accounts/v1") {
		ptHandleRequest(w, path)
	} else {
		w.WriteHeader(400)
		fmt.Fprint(w, "Page Not Found!")
	}
}

func ptHandleRequest(w http.ResponseWriter, path string) {
	if err == nil {
		switch path {
		case "/pt-mock/accounts/v1":
			fmt.Fprint(w, string(ptGetAccounts))
		}
	} else {
		log.Fatal(err)
	}

}
