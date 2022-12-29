package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	file, err := os.Open("tickets.csv")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Hubo un error durante la ejecución. Error: %s\n", err)
		}
		file.Close()
	}()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	// EJERCICIO 3
	response, err := tickets.AverageDestination("United States", data)
	if err != nil {
		panic(err)
	}

	fmt.Println(response, "%")
}
