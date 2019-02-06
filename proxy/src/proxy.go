package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func Forward(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	url := fmt.Sprintf("http://webdav-apache/webdav/%s", r.URL.Path)
	req, _ := http.NewRequest(r.Method, url, r.Body)
	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	for key, value := range r.Header {
		req.Header.Set(key, strings.Join(value, ","))
	}
	resp, _ := client.Do(req)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	for key, value := range resp.Header {
		w.Header().Set(key, strings.Join(value, ","))
	}
	fmt.Fprintf(w, string(bodyBytes))
}

func main() {
	http.HandleFunc("/", Forward)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
