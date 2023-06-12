package scanner

import "fmt"

func ShowHelp() {
	head := "Backdoor Scanner Tool"
	tail := "Options: -h, --help Show this help menu"

	fmt.Printf("%s \n --------------- \n Usage: go run backdoor_scanner.go \n %s", head, tail)
}