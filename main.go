// Package static demonstrates a static file handler for App Engine flexible environment.
package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/home/", homepageHandler)
	http.Handle("/index.html", http.FileServer(http.Dir("dist")))
	http.Handle("/", http.FileServer(http.Dir("dist")))
	appengine.Main()
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("Hello world"))
}
