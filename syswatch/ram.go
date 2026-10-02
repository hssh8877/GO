package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMemInfo() {
	file, err := os.Open("/proc/meminfo")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	var total uint64
	var available uint64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal: ") {
			fields := strings.Fields(line)
			total, err = strconv.ParseUint(fields[1], 10, 64)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		if strings.HasPrefix(line, "MemAvailable") {
			fields := strings.Fields(line)
			available, err = strconv.ParseUint(fields[1], 10, 64)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	used := total - available
	availableGB := float64(available) / 1024 / 1024
	totalGB := float64(total) / 1024 / 1024
	usedGB := float64(used) / 1024 / 1024
	usagePercent := usedGB / totalGB * 100

	fmt.Printf("Total: %.2f GB\n", totalGB)
	fmt.Printf("Available: %.2f GB\n", availableGB)
	fmt.Printf("Used: %.2f GB\n", usedGB)
	fmt.Printf("Usage: %.2f%%\n", usagePercent)
}
