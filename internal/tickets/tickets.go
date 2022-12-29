package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Horarios struct {
	Madrugada int
	Mañana    int
	Tarde     int
	Noche     int
}

var (
	InternalServerError = fmt.Errorf("Hubo un error en el sistema, por favor intente más tarde")
	ErrEmptyDestination = fmt.Errorf("El destino está vacío. Escriba uno.")
)

type Ticket struct {
}

// ejemplo 1
func GetTotalTickets(destination string) (response int, err error) {
	if destination == "" {
		err = ErrEmptyDestination
	}

	data, err1 := getClientsInformation()
	if err1 != nil {
		err = err1
		return
	}

	for _, client := range data {
		if client[3] == destination {
			response++
		}
	}
	return
}

// ejemplo 2
func GetCountByPeriod() (response interface{}, err error) {
	data, err1 := getClientsInformation()
	if err1 != nil {
		err = err1
		return
	}

	horarios := &Horarios{}
	for _, client := range data {
		clientHorario := strings.Split(client[4], ":")[0]
		err1 := addInformationByHour(horarios, clientHorario)
		if err != nil {
			err = err1
		}
	}
	response = *horarios
	return
}

func addInformationByHour(horarios *Horarios, hora string) (err error) {
	n, err := strconv.Atoi(hora)
	if err != nil {
		err = InternalServerError
	}

	if n >= 0 && n < 6 {
		horarios.Madrugada++
	} else if n >= 6 && n < 12 {
		horarios.Mañana++
	} else if n >= 12 && n < 19 {
		horarios.Tarde++
	} else {
		horarios.Noche++
	}
	return
}

// ejemplo 3
func AverageDestination(destination string) (response int, err error) {
	totalTickets, err1 := GetTotalTickets(destination)
	if err1 != nil {
		err = err1
		return
	}

	data, err1 := getClientsInformation()
	if err1 != nil {
		err = err1
		return
	}

	response = totalTickets * 100 / len(data)
	return
}

func getClientsInformation() (data [][]string, err error) {
	file, err := os.Open("../../tickets.csv")
	if err != nil {
		err = InternalServerError
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Hubo un error durante la ejecución. Error: %s\n", err)
		}
		file.Close()
	}()

	csvReader := csv.NewReader(file)
	response, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	data = response
	return

}
