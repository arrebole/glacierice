package controller

import (
	"encoding/json"
	"net/http"

	"github.com/arrebole/glacierice/service"
)

// Response ...
func Response(code int, msg string, data interface{}) []byte {
	result, _ := json.Marshal(
		map[string]interface{}{
			"code":    code,
			"message": msg,
			"data":    data,
		})
	return result
}

// DataCtl ...
func DataCtl() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.FormValue("type") {
		case "contributor":
			w.Write(Response(0, "success", service.Contributor()))
		case "resource_github":
			w.Write(Response(0, "success", service.ResourceGithub()))
		case "resource_3rd":
			w.Write(Response(0, "success", service.Resource3RD()))
		default:
			w.Write(Response(-1, "bad query type", nil))
		}

	}
}
