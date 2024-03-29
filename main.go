package main

import (
	"fmt"
	"net/http"
	"os"
)

var URL string

func main() {
	URL = parseFlags()

	_, err := os.Stat("./hosted")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("./hosted", 0755)
			err_panic(err)
		}
	}
	// Base URL, redirect to template site
	http.Handle("/", http.FileServer(http.Dir("./site")))

	// Upload function, endpoint to post files to the server (POST method)
	http.HandleFunc("/upload", upload)

	// Hosted function, endpoint to get files from the server (GET method)
	http.HandleFunc("/hosted/", serve)

	fmt.Println("Server started and live at " + URL)

	// Listen to localhost, change according your needs
	err = http.ListenAndServe(URL, nil)
	err_panic(err)
}
