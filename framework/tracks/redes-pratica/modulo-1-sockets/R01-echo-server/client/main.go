package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	launchClients(&wg)
	wg.Wait()
}

func stablish(wg *sync.WaitGroup) {
	// estabeece a conexão tcp no endereço indicado
	defer wg.Done()
	conn, err := net.Dial("tcp", "localhost:9000")
	defer conn.Close()

	// buffer para ler o que vier da conexão(limitado a 1kb)
	buf := make([]byte, 1024)
	if err != nil {
		//não pode ser fatal porque fatal panica, mas não temos como ler uma conexão que é nil, tampouco escrever
		log.Println("Error while stablishing a client")
		return
	}
	// precisamos garantir que a conexão seja fechada após a mensagem ser enviada

	//escrevemos na conexão um array de bytes
	conn.Write([]byte("Hello"))

	// lemos o que vier e despejamos no buffer
	n, _ := conn.Read(buf)

	// printamos o dado, limitado ao tamanho N do buffer
	fmt.Println(string(buf[:n]))
}

// lanço 100 conexões
func launchClients(wg *sync.WaitGroup) {
	for _ = range 100 {
		wg.Add(1)
		go stablish(wg)
	}
}
