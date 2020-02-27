package main

import (
	"net/http"

	"github.com/arrebole/glacierice/config"
	"github.com/arrebole/glacierice/controller"
)

func main() {
	http.HandleFunc("/", controller.FileCtl())
	http.HandleFunc("/api/data", controller.DataCtl())
	http.ListenAndServe(config.ListenHost+":"+config.ListenPort, nil)
}
