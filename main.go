package main

import (
	"HttpLogger/HttpLog"
	"flag"
)

func main() {
	portNumber := flag.String("p", "5123", "Port number to listen on")
	flag.Parse()

	HttpLog.Log(*portNumber)
}
