package main

import (
	"fmt"
	"syscall"
)

func readDiskInfo() {
	var stat syscall.Statfs_t

	err := syscall.Statfs("/", &stat)
	if err != nil {

		return
	}
	totalBytes := stat.Blocks * uint64(stat.Bsize)
	totalGB := float64(totalBytes) / 1024 / 1024 / 1024
	availableBytes := stat.Bavail * uint64(stat.Bsize)
	availableGB := float64(availableBytes) / 1024 / 1024 / 1024
	usedGB := totalGB - availableGB
	usage := usedGB / totalGB * 100

	fmt.Printf("Total: %.2f GB\n", totalGB)
	fmt.Printf("Available: %.2f GB\n", availableGB)
	fmt.Printf("Used: %.2f GB\n", usedGB)
	fmt.Printf("Usage: %.2f%%\n", usage)
}
