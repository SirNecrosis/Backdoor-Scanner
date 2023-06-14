package scanner

import (
	"net"
	"strconv"
	
	"github.com/SirNecrosis/Backdoor-Scanner/tools"
	"github.com/SirNecrosis/Backdoor-Scanner/types"
)

func Conf() types.Config {
	// Extend default port support
	// ref: https://nmap.org/book/port-scanning.html#most-popular-ports
	p := []int{
		21, 
		22, 
		23, 
		25, 
		53,  
		68,
		80,
		110, 
		111,
		135,
		137, 
		443,
		445, 
		514, 
		631, 
		993, 
		1723, 
		3306, 
		3389, 
		5900, 
		8080,
	}
	return types.Config{
		BackdoorPorts: p,
		Protocol:      []string{"TCP", "UDP"},
	}
}

func PortOpt(ip string, port int, protocol string) int {
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
	} 
	
	if err != nil {
		tools.Log().Fatal("Failed to enumerate ports!")
	}

	defer sock.Close()
	return 0
}