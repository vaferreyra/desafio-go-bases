package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// EJERCICIO 3
	response, err := tickets.AverageDestination("United States")
	if err != nil {
		panic(err)
	}

	fmt.Println(response, "%")
}
