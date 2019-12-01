package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	listenPort int
	uploadPath string
)

func init() {
	// check for user specified log level and set default if not found
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
	log.SetOutput(os.Stdout) // log output stdout - see 12 factor app design

	// Set and parse flags for configuration with sane defaults
	flag.IntVar(&listenPort, "port", 3500, "web server listening port")
	flag.IntVar(&listenPort, "p", 3500, "short flag for web server listening port")
	flag.StringVar(&uploadPath, "upload-path", os.TempDir(), "Directory to write uploaded files to")
	flag.StringVar(&uploadPath, "u", os.TempDir(), "Directory to write uploaded files to")
	flag.Parse()

	// summary of config after init - for testing
	fmt.Printf("Application initialised successfully:\nweb server port:\t %v\nupload directory:\t %v\nlog level:\t\t %v\n", listenPort, uploadPath, log.GetLevel())
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s", r.URL.Path[1:])
}

func upload() {
	fmt.Println("This function will handle uploading files")
}

func main() {
	// initial configuration (logging, ports, etc.)
	// run web server
	upload()
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
	if err != nil {
		log.Error(err)
	}
}
