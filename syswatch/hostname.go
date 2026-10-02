package main

import (
	"os"
)

func getHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		return "ERROR: Hostname not found!"
	}

	return hostname
}
