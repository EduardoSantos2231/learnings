package main

import "fmt"

func main() {

	original := make([]string, 0, 5)
	original = append(original, "a", "b", "c", "d")

	unsafePop, deletedVal := pop(original)
	fmt.Println("------POP SEM DELETAR VALOR------")
	fmt.Println("Provando que o elemento ainda é acessível caso não zerado: ", unsafePop[:cap(unsafePop)])
	fmt.Println("Valor a ter sido deletado: ", deletedVal)
	safePop, previousLastVal := popSafe(original)
	fmt.Println("------POP COM DELETAR VALOR------")
	fmt.Println("Provando que o valor agora foi zerado: ", safePop[:cap(safePop)])
	fmt.Println("Valor que existia antes: ", previousLastVal)

	fmt.Println("-------DEQUEUE DELETANDO VALOR-------")
	originalDequeued, deleted := dequeue(original)
	fmt.Println("Novo primeiro elemento: ", originalDequeued[0])
	fmt.Println("Valor deletado: ", deleted)
	fmt.Println("Tentando acessar valor antigo: ", originalDequeued[:cap(originalDequeued)])
	originalDequeued = append(originalDequeued, "teste")
	fmt.Println("Fazendo append: ", originalDequeued[:cap(originalDequeued)])
}

func pop(s []string) ([]string, string) {

	// pegar ultimo indice
	lastIndex := len(s) - 1

	//retornar o slice do inicio até o último index (não inclusive) e o elemento que foi apagado
	return s[:lastIndex], s[lastIndex]
}

func popSafe(s []string) ([]string, string) {
	lastIndex := len(s) - 1
	lastVal := s[lastIndex]
	//zeramos o último elemento antes de fazer o fatiamento
	s[lastIndex] = ""

	return s[:lastIndex], lastVal
}

func dequeue(s []string) ([]string, string) {
	firstEl := s[0]
	s[0] = ""
	return s[1:], firstEl
}
