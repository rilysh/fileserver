package main

import (
	"flag"
	"strconv"
)

// Parse specific flags for domain and port
func parseFlags() string {
	address := flag.String("d", "127.0.0.1", "IP/Domain address")
	port := flag.Int("p", 1337, "Port number")

	flag.Parse()

	return *address + ":" + strconv.Itoa(*port)
}
