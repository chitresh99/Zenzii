package main

import (
	"flag"
)

func main() {
	mode := flag.String("mode", "server", "Choose 'server' or 'client'")
	flag.Parse()

	if *mode == "server" {
		server()
	} else if *mode == "client" {
		client()
	} else {
		panic("Unknown mode. Use -mode=server or -mode=client")
	}
}
