package main

import "fmt"

func suma(a, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}

func ResultadoSuma(a int, b int) {
	fmt.Println("la suma de", a, "y", b, "es; ", suma(a, b))
}
func ResultadoResta(a int, b int) {
	fmt.Println("la resta de", a, "y", b, "es; ", resta(a, b))
}

func main() {

	//var a int = 5
	//var b int = 8

	//fmt.Println("la suma de a + b es; ", suma(a, b))

	//fmt.Println("la suma de 2 + 3 es; ", suma(2, 3))

	ResultadoSuma(10, 2)

	ResultadoResta(10, 2)

	ResultadoSuma(23, 5)

	ResultadoResta(23, 5)

	ResultadoSuma(1, 9)

	ResultadoResta(1, 9)
}
