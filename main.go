package main

// todo: webserver listening on listenPort
// todo: gohtml templates for front end / form data
// todo: function to write data from form data to file at uploadPath

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {

	listenPort := flag.Int("port", 3500, "Port number to serve on")
	uploadPath := flag.String("upload-path", "/tmp", "Upload path of files")

	portString := strconv.Itoa(*listenPort)

	uploadPathInt := *uploadPath

	fmt.Println("The web server will listen on port:" + portString + " and upload files to " + uploadPathInt)

}
