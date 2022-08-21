package main

import (
	"io/fs"
	"path/filepath"
)

// Clean hosted directory when needed
func clean() {
	for {
		filepath.Walk("./hosted", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			mb := info.Size() / (1024 * 1024)
			switch mb {
			case 10:
				// 10 minutes
				go sleep_delete(10, path)

			case 9:
				// 15 minutes
				go sleep_delete(15, path)

			case 8:
				// 25 minutes
				go sleep_delete(25, path)

			case 7:
				// 35 minutes
				go sleep_delete(35, path)

			case 6, 5:
				// Around 5 to 6, 45 minutes
				go sleep_delete(45, path)

			case 4, 3:
				// Around 3 to 5, 55 minutes
				go sleep_delete(55, path)

			case 2, 1, 0:
				// Around 0 to 2, 60 minutes (note 0 here refers 0 byte size of the file, empty)
				go sleep_delete(60, path)
			}
			return err
		})
	}
}
