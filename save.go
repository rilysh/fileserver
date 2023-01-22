package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

// Create and save the file in the "hosted" folder
func save(filename string, part *multipart.FileHeader) (string, error) {
	ext := strings.Split(filename, ".")
	last_val := ext[len(ext)-1]
	gen_file := "./hosted/" + strconv.FormatInt(assign_rand(), 36) + "." + last_val

	file, err := os.Create(gen_file)
	err_panic(err)

	file.Sync()
	defer file.Close()

	source, err := part.Open()
	err_panic(err)

	_, err = io.Copy(file, source)
	err_panic(err)

	// Wait for 3 secoonds until the io.Copy finishes to write to the disk
	// And filesystem synchronize everything (if needed).
	time.Sleep(time.Second * 3)

	raw, err := os.ReadFile(file.Name())
	err_panic(err)

	encRaw := encrypt(string(raw), []byte(authKey()))
	if err := os.Remove(file.Name()); err != nil {
		panic(err)
	}

	newFile, err := os.Create(file.Name())
	err_panic(err)

	_, err = io.Copy(newFile, bytes.NewReader([]byte(encRaw)))
	err_panic(err)

	return strings.ReplaceAll(file.Name(), "./hosted/", "/hosted/"), err
}
