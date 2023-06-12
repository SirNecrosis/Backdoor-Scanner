package scanner

import (
	"net"
	"strconv"

	"github.com/SirNecrosis/Backdoor-Scanner/tools"
)

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