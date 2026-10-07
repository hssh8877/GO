package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func calcCpuUsage() (uint64, uint64, uint64) {
	file, err := os.Open("/proc/stat")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0, 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)

			user, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return 0, 0, 0
			}

			system, err := strconv.ParseUint(fields[3], 10, 64)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return 0, 0, 0
			}

			idle, err := strconv.ParseUint(fields[4], 10, 64)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return 0, 0, 0
			}

			return user, system, idle
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0, 0, 0
	}

	return 0, 0, 0
}

func readCPUDetails() {
	file, err := os.Open("/proc/cpuinfo")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	vendorfound := false
	modelfound := false
	cpucoresfound := false
	threadsFound := false

	frequencies := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if !vendorfound && strings.HasPrefix(line, "vendor_id") {
			fields := strings.Fields(line)
			fmt.Println("Vendor: ", fields[2])
			vendorfound = true
		}

		if !modelfound && strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			model := strings.TrimSpace(parts[1])
			fmt.Println("Model: ", model)
			modelfound = true
		}

		if !cpucoresfound && strings.HasPrefix(line, "cpu cores") {
			fields := strings.Fields(line)
			fmt.Println("CPU cores: ", fields[3])
			cpucoresfound = true
		}

		if !threadsFound && strings.HasPrefix(line, "siblings") {
			fields := strings.Fields(line)
			fmt.Println("Threads: ", fields[2])
			threadsFound = true
		}

		if strings.HasPrefix(line, "cpu MHz") {
			fields := strings.Fields(line)
			frequencies = append(frequencies, fields[3])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	arch := runtime.GOARCH
	fmt.Printf("Architecture: %s\n", arch)

	fmt.Println("CPU Frequencies")

	for i, frequency := range frequencies {
		fmt.Printf("CPU %d: %s MHz\n", i, frequency)
	}
}
