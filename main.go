package main

import (
	"HttpLogger/HttpLog"
	"flag"
)

func main() {

	portNumber := flag.String("p", "5123", "Port number to listen on")
	logDir := flag.String("o", "./", "Location of log output")
	//detached := flag.Bool("d", false, "Detact and run in background")

	flag.Parse()

	HttpLog.Log(HttpLog.ConfigurationOptions{
		Port:          *portNumber,
		LogDir:        *logDir,
		RunAsDetached: false,
	})
}
