package main

import "fmt"

var funciones = map[string]func(int, int) int{
	"suma":           func(a, b int) int { return a + b },
	"resta":          func(a, b int) int { return a - b },
	"multiplicacion": func(a, b int) int { return a * b },
	"division":       func(a, b int) int { return a / b },
}

func MostrarResultado(operacion string, a int, b int) {

	f, exists := funciones[operacion]
	if !exists {
		fmt.Println("operacion no valida")
		return
	}

	fmt.Println("para a = ", a, "y b = ", b, "la ", operacion, "resultante es: ", f(a, b))
}

func main() {

	MostrarResultado("suma", 10, 2)
	MostrarResultado("resta", 10, 2)
	MostrarResultado("division", 10, 2)
	MostrarResultado("", 10, 2)
	MostrarResultado("suma", 23, 5)
	MostrarResultado("multiplicacion", 23, 5)
}
