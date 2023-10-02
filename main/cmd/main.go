package main

import (
	"fmt"
	"myapp/pkg/network"
)

func main() {
	devices, err := network.DiscoverDevices()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, device := range devices {
		fmt.Println(device)
	}
	fmt.Println("Done Network Scanning, beginning Port Scanning")
	network.ScanPorts()
}
