package scanner

import "fmt"

func RunScan(ip string, port int, protocol string) {
	if PortOpt(ip, port, protocol) == 0 {
		fmt.Printf("Port %d open on %s. Possible backdoor detected.\n", port, ip)
	} 
	fmt.Printf("Port %d closed on %s. No backdoor detected.\n", port, ip)	
}