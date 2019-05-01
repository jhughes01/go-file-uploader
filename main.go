package main

/*
done: webserver listening on listenPort
todo: gohtml templates for front end / form data
todo: function to write data from form data to file at uploadPath
*/

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)

	if r.Method == "GET" {

	}
}

func main() {

	listenPort := flag.Int("port", 3500, "Port number to serve on")
	uploadPath := flag.String("upload-path", "/tmp", "Upload path of files")

	portString := strconv.Itoa(*listenPort)

	uploadPathInt := *uploadPath

	fmt.Println("The web server will listen on port:" + portString + " and upload files to " + uploadPathInt)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	// Super simple Health check - returns 200 OK if up
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	// Upload function

	http.HandleFunc("/upload", upload)

	http.ListenAndServe(":"+portString, nil)

}
