package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	// EJERCICIO 1
	number, err := tickets.GetTotalTickets("Argentina")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", number)
}
