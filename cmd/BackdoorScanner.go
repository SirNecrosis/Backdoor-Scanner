package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	//"reflect"
	// "flag"

	"github.com/SirNecrosis/Backdoor-Scanner/tools"
)

type Config struct {
	backDoorPorts []int
	protocol      []string
}

// backdoorPorts := []int{21, 22, 23, 25, 53, 80, 443, 3306, 3389, 8080}
func ScannerConfig() Config {
	return Config{
		backDoorPorts: []int{21, 22, 23, 25, 53, 80, 443, 3306, 3389, 8080},
		protocol:      []string{"TCP", "UDP"},
	}
}

func checkPort(ip string, port int, protocol string) int {
	var sock net.Conn
	var err error

	if protocol == ScannerConfig().protocol[0] {
		sock, err = net.Dial(strings.ToLower(ScannerConfig().protocol[0]), ip+":"+strconv.Itoa(port))
	} else if protocol == ScannerConfig().protocol[1] {
		sock, err = net.Dial(strings.ToLower(ScannerConfig().protocol[1]), ip+":"+strconv.Itoa(port))
	} else {
		sock, err = net.Dial(strings.ToLower(ScannerConfig().protocol[0]), ip+":"+strconv.Itoa(port))
	}

	if err != nil {
		tools.Log().Fatal("Failed to enumerate ports!")
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
	fmt.Printf("Backdoor Scanner Tool \n --------------- \n Usage: go run backdoor_scanner.go \n Options: -h, --help Show this help menu")
}

// func GetPort() - Grabs the prot from the user
// TODO if there is a network error, it returns an error
// func GetPort(port int32) string {
// Get port selection from user
// var port int32
// fmt.Println("Enter a port to enumerate...")

// captured user input
// err, x := fmt.Scan(&port)

// }

// func GetProtocol
// func GetProtocol(protocol string) (string, net.Error) {

// }

//func (c *Config) AddProtocol() {

// }

//func (c *Config) AddPort() {

// }

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
	// Prompt user to input target IP address
	var targetIP string
	fmt.Print("Enter the IP address to scan for backdoors: ")
	fmt.Scan(&targetIP)

	// Choose a protocol (More to come in the near future!)
	var protocol string
	fmt.Print("Enter the protocol to scan for (TCP/UDP/Both): ")
	fmt.Scan(&protocol)

	for _, port := range ScannerConfig().backDoorPorts {
		scanBackdoor(targetIP, port, protocol)
	}
}
