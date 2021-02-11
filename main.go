package main

import (
	"fmt"
	"net"
)

func main() {
	_, erro := net.Dial("tcp", "scanme.nmap.org:80")
		if erro == nil {
			fmt.Println("Conectado com sucesso.")
		}
}