package main

import (
	"net/http"
	"strings"
)

// Serve the file when requested
func serve(w http.ResponseWriter, r *http.Request) {
	if len(strings.Split(r.URL.Path, "/")[2]) == 0 {
		w.Write([]byte("Missing file name"))
		return
	}

	files := http.StripPrefix("/hosted/", http.FileServer(http.Dir("./hosted")))
	files.ServeHTTP(w, r)
}
