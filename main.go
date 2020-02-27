package main

import (
	"flag"
	"net/http"

	"github.com/arrebole/glacierice/controller"
)

var (
	listenHost string
	listenPort string
)

func init() {
	flag.StringVar(&listenHost, "host", "127.0.0.1", "http listen host")
	flag.StringVar(&listenPort, "port", "3000", "http listen port")
	flag.Parse()
}

func main() {

	http.HandleFunc("/", controller.FileCtl())
	http.HandleFunc("/api/data", controller.DataCtl())
	http.ListenAndServe(listenHost+":"+listenPort, nil)
}
