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
	fmt.Println("Blocks: ", stat.Blocks)
	fmt.Println("Block size: ", stat.Bsize)
	fmt.Println("Total bytes: ", totalBytes)
}
