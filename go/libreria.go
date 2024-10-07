package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Alumno struct {
	Nombre string
	Notas  []float64
}

func leerCSV(nombreArchivo string) ([]Alumno, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}

	lector := csv.NewReader(archivo)
	lector.Comma = ';'

	registros, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}
	var alumnos []Alumno

	for _, registro := range registros {
		nombre := registro[0]
		var notas []float64
		for _, notaStr := range registro[1:] {
			nota, err := strconv.ParseFloat(notaStr, 64)
			if err != nil {
				return nil, err
			}
			notas = append(notas, nota)
		}
		alumnos = append(alumnos, Alumno{Nombre: nombre, Notas: notas})
	}
	return alumnos, nil
}

func CalcularPromedio(notas []float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}
func CalcularMediaArt(alumnos []Alumno) float64 {
	var suma float64
	var cantidadNotas int
	for _, alumno := range alumnos {
		for _, nota := range alumno.Notas {
			suma += nota
			cantidadNotas++
		}

	}
	return suma / float64(cantidadNotas)
}

func EscribirCSV(nombreArchivo string, alumnos []Alumno) error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	escritor := csv.NewWriter(archivo)
	defer escritor.Flush()
	for _, alumno := range alumnos {
		promedio := CalcularPromedio(alumno.Notas)
		registro := []string{alumno.Nombre, fmt.Sprintf("%.2f", promedio)}
		if err := escritor.Write(registro); err != nil {
			return err
		}
	}
	return nil

}
func main() {

	alumnos, err := leerCSV("notas.csv")
	if err != nil {
		fmt.Println("error al leer archivo", err)
		return
	}
	mediaArt := CalcularMediaArt(alumnos)
	fmt.Println("La media de todas las notas es: %.2f\n	", mediaArt)

	err = EscribirCSV("promedios.csv", alumnos)
	if err != nil {
		fmt.Println("Eror al ascribir archivo", err)
		return
	}
	fmt.Println("Archivo escrito con exito")

}
