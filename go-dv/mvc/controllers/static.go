package controllers

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func Static(staticDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// restrict requests to "GET"
		if r.Method != "GET" {
			log.Printf("METHOD ERROR: method not allowed")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		path := r.URL.Path

		// ensure only allowed static directories are served
		allowedPrefixes := []string{"/static/", "/css/", "/js/", "/images/"}
		valid := false
		var strippedPath string

		for _, prefix := range allowedPrefixes {
			if strings.HasPrefix(path, prefix) {
				valid = true
				// Strip prefix if it's "/static/"
				if prefix == "/static/" {
					strippedPath = strings.TrimPrefix(path, "/static/")
				} else {
					strippedPath = path[1:] // remove leading "/"
				}
				break
			}
		}

		if !valid {
			log.Println("DIRECTORY AVAILABILITY ERROR: directory not found")
			http.NotFound(w, r)
			return
		}

		// Construct full filesystem path
		filePath := staticDir + "/" + strippedPath
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Printf("FILE AVAILABILITY ERROR: %v", err)
			http.NotFound(w, r)
			return
		}

		// Restrict access to directories
		if fileInfo.IsDir() {
			log.Println("DIRECTORY AVAILABILITY ERROR: directory not found")
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, filePath)
	}
}
