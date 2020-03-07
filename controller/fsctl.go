package controller

import (
	"net/http"
	"os"
	"path/filepath"
)

// FileCtl ...
func FileCtl() http.HandlerFunc {
	var StaticRoot, _ = filepath.Abs("./theme/dist")

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")

		if fileExist(filepath.Join(StaticRoot, r.RequestURI)) {
			http.FileServer(http.Dir(StaticRoot)).ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, filepath.Join(StaticRoot, "index.html"))
	}
}

// fileExist 判断文件是否存在
func fileExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
