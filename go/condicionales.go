package main

import (
	"fmt"
)

func condicionales() {
	var edad int = 21
	var permiso bool = false

	/* if edad >= 65 {
		fmt.Print("la persona puede estar jubilada")
	} else if edad >= 18 {
		fmt.Print("la persona es economicamente activa")
	} else {
		fmt.Print("la persona es menor de edad")
	}
	fmt.Print("fin del programa") */

	if edad < 18 && permiso {
		fmt.Print("el menor de edad puede viajar solo")
	} else if edad < 18 && !permiso {
		fmt.Print("el menor de edad NO puede viajar solo")
	} else {
		fmt.Print("la persona puede viajar sola")
	}
	fmt.Print("fin del programa")

}
