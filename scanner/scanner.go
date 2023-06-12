package scanner

import (
	"fmt"

	"github.com/SirNecrosis/Backdoor-Scanner/types"
)

func Conf() types.Config {
	return types.Config{
		BackdoorPorts: []int{21, 22, 23, 25, 53, 80, 443, 3306, 3389, 8080},
		Protocol:      []string{"TCP", "UDP"},
	}
}

func RunScan(ip string, port int, protocol string) {
	if PortOpt(ip, port, protocol) == 0 {
		fmt.Printf("Port %d open on %s. Possible backdoor detected.\n", port, ip)
	} 
	fmt.Printf("Port %d closed on %s. No backdoor detected.\n", port, ip)	
}