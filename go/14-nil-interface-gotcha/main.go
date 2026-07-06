package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
)

func main() {
	defer explain()
	// o tipo é uma interface que implementa o método Write
	// por não ser inicializado ele recebe nil
	var w io.Writer

	// buf é um pointer para um struct que implementa o método Write
	// buf tem valor nil, pois não é inicializado
	var buf *bytes.Buffer

	// o valor de w de fato é nil aqui, pois o zero-value de uma interface é nil
	fmt.Println("w == nil: ", w == nil)

	// w passa a ser buf, que possui um ponteiro para um buffer
	// buffer possui sim o método Write, mas note que ele não foi inicializado, ele é um pointer, e o zero-value de um pointer é nil
	w = buf

	//w agora guarda o valor de buff (um pointer) e a saída é false
	fmt.Println("w == nil: ", w == nil)

	fmt.Println("About to call w.Write...")

	// ao tentar chamar o método nós temos um problema...como vamos fazer o derefferecing de um ponteiro que não aponta pra lugar nenhum?
	w.Write([]byte("hello"))
}

// função que mostra o erro antes de a execução crashar
func explain() {
	r := recover()
	fmt.Printf("\nPANIC: ")
	fmt.Println(r)
}

func safeWrite(w io.Writer, data []byte) (int, error) {
	if reflect.ValueOf(w).IsNil() {
		return 0, errors.New("Invalid writer was provided")
	}
	return w.Write(data)
}
