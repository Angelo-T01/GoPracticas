package main

import (
	"fmt"
)

func condicion() {
	var fruta string

	fmt.Print("ingresa un valor:")
	fmt.Scan(&fruta)

	switch fruta {
	case "manzana":
		fmt.Print("has ingresado manzana .")
	case "banana":
		fmt.Print("has ingresado banana .")
	case "pera":
		fmt.Print("has ingresado pera .")
	default:
		fmt.Print("no se ha reconocido el valor")
	}
}
