package middlewares

import (
	"net/http"
	"path/filepath"
	"strings"
)

func ServeStatic(urlPath string, fsPath string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if urlPath[0] != '/' {
			panic("invalid urlPath: " + urlPath)
		}
		relPathToUrlPath, err := filepath.Rel(urlPath, r.URL.Path)
		if err != nil || strings.Contains(relPathToUrlPath, "..") || r.Method != "GET" {
			h.ServeHTTP(w, r)
			return
		}
		absFsPath, err := filepath.Abs(fsPath)
		if err != nil {
			panic(err)
		}
		absReqPath := filepath.Join(absFsPath, relPathToUrlPath)
		http.ServeFile(w, r, absReqPath)
	})
}
