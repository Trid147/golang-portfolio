package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"
)

func CheckPort(ip string, port int) bool {
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	timeout := 2 * time.Second

	connection, error := net.DialTimeout("tcp", address, timeout)
	if error != nil {
		return false
	}
	connection.Close()
	return true
}

func main() {
	// CONFIGURATION
	var IP = "192.168.0.1"                                       // you can change to 192.168.1.1 if your router use it instead
	ports := []int{21, 22, 23, 25, 80, 110, 139, 443, 445, 3389} // you can add or remove any ports

	portChan := make(chan int, len(ports))
	var wg sync.WaitGroup

	fmt.Println("Scanning started...")
	startTime := time.Now()

	for _, port := range ports {
		wg.Add(1)

		go func(p int) {
			defer wg.Done()

			if CheckPort(IP, p) {
				portChan <- p
			}
		}(port)
	}

	wg.Wait()

	close(portChan)

	var resultList []int
	for port := range portChan {
		resultList = append(resultList, port)
	}

	sort.Ints(resultList)

	fmt.Println("RESULTS:")
	if len(resultList) == 0 {
		fmt.Println("All checked ports are closed.")
	} else {
		for _, port := range resultList {
			fmt.Printf("Port %d open\n", port)
		}
	}

	fmt.Printf("Scan time: %v\n", time.Since(startTime))
}
