package main

import "fmt"

func main() {
	miMapa := map[string]string{
		"Mexico":   "CDMX",
		"Colombia": "Bogota",
		"Chlie":    "Santiago",
	}

	fmt.Println("la capital de mexico es; ", miMapa["Mexico"])

	miMapa["Argentina"] = "Buenos Aires"

	fmt.Println("el mapa de paises ", miMapa)

	delete(miMapa, "Colombia")

	fmt.Println("el mapa de paises ", miMapa)

	fmt.Println("la capital de mexico es; ", miMapa["Colombia"])
}
