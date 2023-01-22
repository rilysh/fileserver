package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

// Serve the file when requested
func serve(w http.ResponseWriter, r *http.Request) {
	if len(strings.Split(r.URL.Path, "/")[2]) == 0 {
		w.Write([]byte("Missing file name"))
		return
	}

	files := http.StripPrefix("/hosted/", http.FileServer(http.Dir("./hosted")))
	file := strings.ReplaceAll(r.URL.Path, "/hosted", "./hosted")

	raw, err := os.ReadFile(file)
	err_panic(err)

	decRaw := decrypt(string(raw), []byte(authKey()))
	if err := os.Remove(file); err != nil {
		panic(err)
	}

	newFile, err := os.Create(file)
	err_panic(err)

	_, err = io.Copy(newFile, bytes.NewReader([]byte(decRaw)))
	err_panic(err)

	files.ServeHTTP(w, r)

	// After serving the requested file, encrypt the requested on the server
	encRaw := encrypt(string(decRaw), []byte(authKey()))
	if err := os.Remove(file); err != nil {
		panic(err)
	}

	newEncFile, err := os.Create(file)
	err_panic(err)

	_, err = io.Copy(newEncFile, bytes.NewReader([]byte(encRaw)))
	err_panic(err)
}
