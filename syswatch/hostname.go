package main

import (
	"fmt"
	"os"
)

func getHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "ERROR: Hostname not found!"
	}

	return hostname
}
