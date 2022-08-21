package main

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

// Create and save the file in hosted folderr
func save(filename string, part *multipart.FileHeader) (string, error) {
	ext := strings.Split(filename, ".")
	last_val := ext[len(ext)-1]
	file, err := os.Create("./hosted/" + strconv.FormatInt(assign_rand(), 36) + "." + last_val)
	if err != nil {
		return "", err
	}
	defer file.Close()

	source, _ := part.Open()
	_, err = io.Copy(file, source)
	return strings.ReplaceAll(file.Name(), "./hosted/", "/hosted/"), err
}
