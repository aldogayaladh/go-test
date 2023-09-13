package main

import (
	"fmt"

	"github.com/aldogayaladh/go-test/funciones"
)

func main() {

	resultado := funciones.Suma(2, 3)
	fmt.Println("El resultado es:", resultado)
	division, err := funciones.Division(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("El resultado de la division es:", division)
}

// Para correr todos los test.
// go test ./...

// Para generar los reportes de todos los test
// go test ./... -coverprofile=cover.out
// go tool cover -html=cover.out

// Para correr un test en un folder especifico
// Ubicados dentro de la carpeta a testear
// go test .

// Para generar los reportes
// Dentro de la carpeta con los test
// go test -coverprofile=cover.out
// go tool cover -html=cover.out
