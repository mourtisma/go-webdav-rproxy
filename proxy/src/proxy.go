package main

import (
	"auth"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Forward sends the request to the Apache webdav server and
// writes the response's headers and body back to the client
func Forward(w http.ResponseWriter, r *http.Request) {

	// Instantiate the client
	client := &http.Client{}

	// Route the request to the webdav server in the webdav-apache container
	url := fmt.Sprintf("http://webdav-apache/webdav/%s", r.URL.Path)

	// Create a new request with the same method and the same body as the one hitting the proxy
	req, _ := http.NewRequest(r.Method, url, r.Body)

	// Inject the basic auth credentials
	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	// Copy the headers
	for key, value := range r.Header {
		req.Header.Set(key, strings.Join(value, ","))
	}

	// Fire the request
	resp, _ := client.Do(req)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Write the status code back
	w.WriteHeader(resp.StatusCode)

	// Write the headers and the response back to the client
	for key, value := range resp.Header {
		w.Header().Set(key, strings.Join(value, ","))
	}

	// Write the body back
	w.Write(bodyBytes)

}

func main() {
	http.HandleFunc("/", Forward)
	http.HandleFunc("/auth/google/login", auth.OauthGoogleLogin)
	http.HandleFunc("/auth/google/callback", auth.OauthGoogleCallback)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
