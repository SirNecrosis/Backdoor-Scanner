package scanner

import (
	"net"
	"fmt"
	"strconv"

	"github.com/SirNecrosis/Backdoor-Scanner/types"
	"github.com/SirNecrosis/Backdoor-Scanner/tools"
)

func Conf() types.Config {
	return types.Config{
		BackdoorPorts: []int{21, 22, 23, 25, 53, 80, 443, 3306, 3389, 8080},
		Protocol:      []string{"TCP", "UDP"},
	}
}

func portOpt(ip string, port int, protocol string) int {
	var (
		sock net.Conn
		err error
	)

	switch protocol {
		case "TCP":
			sock, err = net.Dial("tcp", ip+":"+strconv.Itoa(port))
		case "UDP":
			sock, err = net.Dial("udp", ip+":"+strconv.Itoa(port))
		default:
			sock, err = net.Dial("tcp", ip+":"+strconv.Itoa(port))
	}; if err != nil {
		tools.Log().Fatal("Failed to enumerate ports!")
	}

	defer sock.Close()
	return 0
}

func RunScan(ip string, port int, protocol string) {
	if portOpt(ip, port, protocol) == 0 {
		fmt.Printf("Port %d open on %s. Possible backdoor detected.\n", port, ip)
	} 
	fmt.Printf("Port %d closed on %s. No backdoor detected.\n", port, ip)	
}

func ShowHelp() {
	fmt.Printf("Backdoor Scanner Tool \n --------------- \n Usage: go run backdoor_scanner.go \n Options: -h, --help Show this help menu")
}