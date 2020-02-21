package main

import (
	"net/http"

	"github.com/arrebole/glacierice/controller"
)

func main() {

	http.HandleFunc("/", controller.FileServer())
	http.ListenAndServe("127.0.0.1:3000", nil)
}
