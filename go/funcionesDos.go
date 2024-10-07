package main

import "fmt"

/* func ResultadoSuma(f func(int, int) int,
	a int, b int) {
	fmt.Println("la suma de", a, "y", b, "es; ", f(a, b))
}

func ResultadoResta(f func(int, int) int,
	a int, b int) {
	fmt.Println("la resta de", a, "y", b, "es; ", f(a, b))
} */

func MostrarResultado(operacion string, a int, b int) {
	suma := func(a, b int) int {
		return a + b
	}

	resta := func(a, b int) int {
		return a - b
	}

	resultado := 0

	if operacion == "sumar" {
		resultado = suma(a, b)
	} else if operacion == "restar" {
		resultado = resta(a, b)
	}

	fmt.Println("para a = ", a, "y b = ", b, "la ", operacion, "resultante es: ", resultado)
}

func main() {

	/* //parametros como funcion
	ResultadoSuma(suma, 10, 2)
	ResultadoResta(resta, 10, 2)
	ResultadoSuma(suma, 23, 5)
	ResultadoResta(resta, 23, 5)
	ResultadoSuma(suma, 1, 9)
	ResultadoResta(resta, 1, 9) */

	MostrarResultado("sumar", 10, 2)
	MostrarResultado("restar", 10, 2)
	MostrarResultado("sumar", 23, 5)
	MostrarResultado("restar", 23, 5)
	MostrarResultado("sumar", 1, 9)
	MostrarResultado("restar", 1, 9)
}
