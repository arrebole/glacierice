package main

import (
	"net/http"

	"github.com/arrebole/glacierice/controller"
)

func main() {

	http.HandleFunc("/", controller.FileServer())
	http.ListenAndServe("0.0.0.0:8080", nil)
}
