package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
		err = InternalServerError
		return
	}

	data = response
	return

}
