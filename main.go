package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(portas, resultados chan int) {
	for p := range portas {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			resultados <- 0
			continue
		}
		conn.Close()
		resultados <- p
	}
}
func main() {
	portas := make(chan int, 100)
	resultados := make(chan int)
	var portasAbertas []int

	for i := 0; i < cap(portas); i++ {
		go worker(portas, resultados)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			portas <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-resultados
		if port != 0 {
			portasAbertas = append(portasAbertas, port)
		}
	}

	close(portas)
	close(resultados)
	sort.Ints(portasAbertas)
	for _, p := range portasAbertas {
		fmt.Printf("porta %d aberta\n", p)
	}
}
