package main

import (
	"fmt"
	"net/http"
)

// Take the file from your computer (client), and save it to the server
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Method not allowed"))
		return
	}

	// Check if content-length is larger than 10 MiB
	if r.ContentLength > 10000001 {
		w.Write([]byte("File size too large"))
		return
	}

	// Multiparse uses memory block, current size 1048576
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		fmt.Println("Failed to allocated memory for incoming form")
		return
	}

	// Can be uploaded more than one file
	files := r.MultipartForm.File["file"]
	if len(files) > 1 {
		w.Write([]byte("Too many files"))
		return
	}
	file, _ := save(files[0].Filename, files[0])
	w.Write([]byte("http://" + URL + file))
}
