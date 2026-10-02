package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=====================")
	fmt.Println("      SYSWATCH")
	fmt.Println("=====================")

	hostname := getHostname()
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Println("-----------------------")
	fmt.Println("Memory")
	readMemInfo()
	fmt.Println("-----------------------")
	fmt.Println("CPU")

	user1, system1, idle1 := calcCpuUsage()
	time.Sleep(1 * time.Second)

	user2, system2, idle2 := calcCpuUsage()
	time.Sleep(1 * time.Second)

	userFinal := user2 - user1
	systemFinal := system2 - system1
	idleFinal := idle2 - idle1

	total := userFinal + systemFinal + idleFinal

	busy := userFinal + systemFinal

	usagePercent := float64(busy) / float64(total) * 100
	fmt.Printf("CPU-Usage: %.2f%%\n", usagePercent)
	readCPUDetails()

	readDiskInfo()

}
