package middlewares

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeStatic(urlPath string, fsPath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(404)
			return
		}
		if urlPath[0] != '/' {
			panic("invalid urlPath: " + urlPath)
		}
		absFsPath, err := filepath.Abs(fsPath)
		if err != nil {
			panic(err)
		}
		relPathToUrlPath, err := filepath.Rel(urlPath, r.URL.Path)
		if err != nil || strings.Contains(relPathToUrlPath, "..") {
			w.WriteHeader(400)
			return
		}
		absReqPath := filepath.Join(absFsPath, relPathToUrlPath)
		f, err := os.Open(absReqPath)
		if err != nil {
			if os.IsNotExist(err) {
				w.WriteHeader(404)
			} else if os.IsPermission(err) {
				w.WriteHeader(403)
			} else {
				w.WriteHeader(400)
			}
			return
		}
		defer f.Close()
		fInfo, err := f.Stat()
		if err != nil {
			// os in Go is too tedious compared to Py...
			w.WriteHeader(400)
			return
		}
		if !fInfo.Mode().IsRegular() {
			// only serve regular file, nginx return
			// 200 by default when requesting directory
			w.WriteHeader(200)
			return
		}
		io.Copy(w, f)
	})
}
