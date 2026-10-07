package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readUptime() {
	file, err := os.Open("/proc/uptime")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		totalSeconds := int64(uptimeSeconds)

		days := totalSeconds / 86400
		remaining := totalSeconds % 86400

		hours := remaining / 3600
		remaining %= 3600

		minutes := remaining / 60
		remaining %= 60

		fmt.Printf("Days: %d\n", days)
		fmt.Printf("Hours: %d\n", hours)
		fmt.Printf("Minutes: %d\n", minutes)
	}

	if err := scanner.Err(); err != nil {
		return
	}
}
