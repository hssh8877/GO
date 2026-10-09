package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("server.log")
	if err != nil {
		fmt.Println("Error opening log file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linecount := 0

	for scanner.Scan() {
		linecount++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading log file: ", err)
		return
	}
	fmt.Println("Total log lines: ", linecount)
}
