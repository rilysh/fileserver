package main

import (
	"flag"
	"strconv"
)

func parse_flags() string {
	address := flag.String("d", "127.0.0.1", "IP/Domain address")
	port := flag.Int("p", 1337, "Port number")

	flag.Parse()

	return *address + ":" + strconv.Itoa(*port)
}