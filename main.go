package main

import (
	"os"
	"fmt"

	"github.com/SirNecrosis/Backdoor-Scanner/scanner"
)

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		scanner.ShowHelp()
		return
	}
	
	if len(os.Args) != 1 {
		fmt.Println("Invalid usage. Use -h or --help for instructions.")
		return
	}
	
	var (
		targetIP string
		protocol string
	)
	// List of ports to scan for potential backdoors
	// Prompt user to input target IP address
	fmt.Print("Enter the IP address to scan for backdoors: ")
	fmt.Scan(&targetIP)
	
	// Choose a protocol (More to come in the near future!)
	fmt.Print("Enter the protocol to scan for (TCP/UDP): ")
	fmt.Scan(&protocol)
	
	for _, port := range scanner.Conf().BackdoorPorts {
		scanner.RunScan(targetIP, port, protocol)
	}
}