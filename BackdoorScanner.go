package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func checkPort(ip string, port int, protocol string) int {
	var sock net.Conn
	var err error

	if protocol == "TCP" {
		sock, err = net.Dial("tcp", ip+":"+strconv.Itoa(port))
	} else if protocol == "UDP" {
		sock, err = net.Dial("udp", ip+":"+strconv.Itoa(port))
	} else {
		sock, err = net.Dial("tcp", ip+":"+strconv.Itoa(port))
	}

	if err != nil {
		return 1
	}
	defer sock.Close()
	return 0
}

func scanBackdoor(ip string, port int, protocol string) {
	if checkPort(ip, port, protocol) == 0 {
		fmt.Printf("Port %d open on %s. Possible backdoor detected.\n", port, ip)
	} else {
		fmt.Printf("Port %d closed on %s. No backdoor detected.\n", port, ip)
	}
}

func showHelp() {
	fmt.Println("Backdoor Scanner Tool")
	fmt.Println("---------------------")
	fmt.Println("Usage: go run backdoor_scanner.go")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help     Show this help menu")
}

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		showHelp()
		return
	}

	if len(os.Args) != 1 {
		fmt.Println("Invalid usage. Use -h or --help for instructions.")
		return
	}

	// List of ports to scan for potential backdoors
	backdoorPorts := []int{21, 22, 23, 25, 53, 80, 443, 3306, 3389, 8080}

	// Prompt user to input target IP address
	var targetIP string
	fmt.Print("Enter the IP address to scan for backdoors: ")
	fmt.Scan(&targetIP)
	var protocol string
	fmt.Print("Enter the protocol to scan for (TCP/UDP/Both): ")
	fmt.Scan(&protocol)

	for _, port := range backdoorPorts {
		scanBackdoor(targetIP, port, protocol)
	}
}
