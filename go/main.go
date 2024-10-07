package main

import "fmt"

func main() {
	//Crear variables
	var nombrePersona string = "Angel"
	var apellidoPersona string = "Ruiz"
	segundoNombre := "obed"

	// numeros
	var asoActual int16 = 2024
	edad := 21

	fmt.Println("hola soy " + nombrePersona)
	fmt.Println(" mi apellido es " + apellidoPersona)
	fmt.Println(" mi segundo nombre es " + segundoNombre)
	fmt.Println(" el año actual es ", asoActual)
	fmt.Println(" mi edad es ", edad)

	// Areglos

	var ListaNombres = [5]string{"Jake", "Eric"}
	fmt.Print(ListaNombres[1])

	//listaPaises := [3]string{"mexico", "peru", "chile"}
	listaPaises := []string{"mexico", "peru", "chile"}
	fmt.Print(listaPaises)
	listaPaises = append(listaPaises, "brasil")
	//listaPaises[0] = "Brasil"

	//fmt.Print(listaPaises)

	// rango
	listaPaises2 := listaPaises[1:3]
	fmt.Print(listaPaises2)
	//Apartir de
	listaPaises3 := listaPaises[2:]
	fmt.Print(listaPaises3)
	//valor no incluido
	listaPaises4 := listaPaises[:2]
	fmt.Print(listaPaises4)
}
