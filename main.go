package main

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
	fmt.Printf("%T", portString)

}
