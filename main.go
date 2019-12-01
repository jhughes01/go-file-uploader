package main

/*
todo: web server listening on listenPort
todo: go html templates for front end / form data
*/

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	listenPort int
	uploadPath string
)

func init() {
	appLogLevel, set := os.LookupEnv("LOG_LEVEL")
	if !set {
		log.SetLevel(log.InfoLevel)
	} else {
		logLevel, err := log.ParseLevel(appLogLevel)
		if err != nil {
			log.Error(err)
		}
		log.SetLevel(logLevel)
	}
	log.SetOutput(os.Stdout)
	flag.IntVar(&listenPort, "port", 3500, "web server listening port")
	flag.IntVar(&listenPort, "p", 3500, "short flag for web server listening port")
	flag.StringVar(&uploadPath, "upload-path", os.TempDir(), "Directory to write uploaded files to")
	flag.StringVar(&uploadPath, "u", os.TempDir(), "Directory to write uploaded files to")
	flag.Parse()
	fmt.Printf("web server port:\t %v\nupload directory:\t %v\nlog level:\t\t %v\n", listenPort, uploadPath, log.GetLevel())
}

func upload() {
	print("This function will handle uploading files")
}

func main() {
	// initial configuration (logging, ports, etc.)
	// run web server
	upload()
}
