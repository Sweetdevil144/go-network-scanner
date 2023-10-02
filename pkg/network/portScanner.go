package network

import (
	"fmt"
	"net"
	"time"
)

func generateGoroutines(host string, ports, resultScanner chan int) {

	for p := range ports {
		address := fmt.Sprintf("%v:%d", host, p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			resultScanner <- 0
			continue
		}
		conn.Close()
		resultScanner <- p
	}
}

func SenderPortsProcess(limitPort int, ports chan int) {
	for i := 1; i <= limitPort; i++ {
		ports <- i
	}
}
func getResultPortsOpen(limitPort int, results chan int) []int {
	var listOpenPorts []int

	for i := 0; i < limitPort; i++ {

		port := <-results

		if port != 0 {

			listOpenPorts = append(listOpenPorts, port)
		}
	}

	return listOpenPorts
}
func ScanPorts() {
	var host string = "scanme.nmap.org"
	var limitPort int = 1024
	ports := make(chan int, limitPort)
	resultsScan := make(chan int)
	for i := 0; i < cap(ports); i++ {
		go generateGoroutines(host, ports, resultsScan)
	}
	go SenderPortsProcess(limitPort, ports)
	portsOpens := getResultPortsOpen(limitPort, resultsScan)
	close(ports)
	close(resultsScan)
	fmt.Printf("Address verification - %v\n", host)
	for _, port := range portsOpens {
		fmt.Printf("%d open\n", port)
	}
}
