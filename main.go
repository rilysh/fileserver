package main

import (
	"fmt"
	"net/http"
)

const URL = "localhost:1337"

func main() {
	// Clean files on hosted directory after a certain amount of time
	go clean()

	// Base URL, redirect to template site
	http.Handle("/", http.FileServer(http.Dir("./site")))

	// Upload function, endpoint to post files to the server (POST method)
	http.HandleFunc("/upload", upload)

	// Hosted function, endpoint to get files from the server (GET method)
	http.HandleFunc("/hosted/", serve)

	fmt.Println("Server started and live at " + URL)

	// Listen to localhost, change according your needs
	http.ListenAndServe(URL, nil)
}
