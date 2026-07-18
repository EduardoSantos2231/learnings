package main

import (
	"fmt"
	"interfaces/shapes"
)

//
// respostas:
// 1) Se declaro um slice vazio com o tipo shape não consigo acessar nenhum elemento para chamar o método Area ou Perimeter
//
// 2) Para satisfazer a interface é necessário implementar os métodos que ela ordena
//
// 3) Não sei

func main() {
	forms := make([]shapes.Shape, 0, 3)
	triangle := shapes.Triangle{
		A: 6.0,
		B: 10,
		C: 8,
	}
	circle := shapes.Circle{Radius: 5.0}
	rectangle := shapes.Rectangle{Width: 8.0, Height: 4.0}
	forms = append(forms, triangle, circle, rectangle)

	for _, format := range forms {

		switch f := format.(type) {
		case shapes.Circle:
			fmt.Println("----Circle----")
			fmt.Println("Perimeter : ", f.Perimeter(), "Area: ", f.Area())
		case shapes.Triangle:
			fmt.Println("----Triangle----")
			fmt.Println("Perimeter : ", f.Perimeter(), "Area: ", f.Area())
		case shapes.Rectangle:
			fmt.Println("----Rectangle----")
			fmt.Println("Perimeter : ", f.Perimeter(), "Area: ", f.Area())
		}
	}

}
