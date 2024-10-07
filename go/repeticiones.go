package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		var suma int = 0

		for i := 0; i <= 100; i++ {
			//if i%2 != 0 {
			suma = suma + i
			//}

		}
		fmt.Println("la suma de los primeros 100 numeros es:", suma)

	*/

	/*
		miMapa := map[string]string{
			"Mexico":   "CDMX",
			"Colombia": "Bogota",
			"Chlie":    "Santiago",
		}


			//FOR EACH

			for k, v := range miMapa {
				fmt.Println(" La capital de " + k + " es:" + v)
			}
	*/
	/* / WHILE
	var fruta string = "pera"
	for {
		if fruta == "naranja" {
			fmt.Println("El usuario finalizo escogiendo naranja")
			break
		}
		fmt.Println("indique su fruta de esperada")
		fmt.Scan(&fruta)
		fruta = strings.ToLower(fruta)
	}
	*/

	//
	var fruta string = "pera"
	for {
		if fruta == "naranja" {
			fmt.Println("El usuario finalizo escogiendo naranja")
			break
		}
		fmt.Println("indique su fruta de esperada")
		fmt.Scan(&fruta)
		fruta = strings.ToLower(fruta)
	}
}
