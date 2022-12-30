package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// EJERCICIO 2
	response, err := tickets.GetCountByPeriod()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", response)
}
