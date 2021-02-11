package main

import (
	"fmt"
	"net"
)

func main() {
    for i := 1; i <= 1024; i++ {
		endereco := fmt.Sprintf("scanme.nmap.org:%d", i)
		conexao, erro := net.Dial("tcp", endereco)

		if erro != nil {
			//porta fechada ou filtrada
			continue
		}

		conexao.Close()
        fmt.Printf("porta %d aberta\n", i)
    }
}