package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var WaitGroup sync.WaitGroup

    for i := 1; i <= 1024; i++ {

		WaitGroup.Add(1)

		go func(j int) {
			
			defer WaitGroup.Done()

			endereco := fmt.Sprintf("scanme.nmap.org:%d", j)
			conexao, erro := net.Dial("tcp", endereco)
	
			if erro != nil {
				//porta fechada ou filtrada
				return
			}
	
			conexao.Close()
			fmt.Printf("porta %d aberta\n", j)
		} (i)
		
		WaitGroup.Wait()
    }
}