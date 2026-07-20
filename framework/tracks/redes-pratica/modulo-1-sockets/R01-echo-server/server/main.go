package main

import (
	"io"
	"log"
	"net"
)

func main() {

	//criação de um listener que ouve na porta 9000 com conexão tcp
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		// se deu ruim na criação do listener não dá pra ouvir nada...
		log.Fatal("\nSomething went wrong: ", err.Error())
	}

	// se não panicou, precisamos garantir que esse listener não vai vazar
	defer listener.Close()

	// vmaos escutar infinitamente
	for {

		// vamos aceitar a conexão
		connection, err := listener.Accept()

		// mesmo com erro, não podemos dar break, isos quebraria a escuta das outras coenxões por vir
		if err != nil {
			log.Println("\nSomething went wrong:", err.Error())
			continue
		}

		// delegamos a uma go routine para lidar com essa conexão
		go handleConnection(connection)
	}

}

func handleConnection(connection net.Conn) {
	//precisamos garantir que essa conexão seja fechada
	defer connection.Close()

	// vamos só copiar o que chegar da conexão
	_, err := io.Copy(connection, connection)

	// se tiver erro não queremos panicar
	if err != nil {
		log.Printf("\nSomething went wrong %s", err.Error())
		return
	}
}
