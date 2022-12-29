package tickets

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// En este test compruebo que el GetTotalTickets devuelve el total de tickets comprados a
// un cierto destino, y el error es nil cuando el destino es correcto
func TestGetTotalTickets_DestinationNotEmpty(t *testing.T) {
	// Arrange
	totalEsperado := 45

	// Act
	totalObtenido, errorObtenido := GetTotalTickets("Brazil")

	// Assert
	assert.Equal(t, totalEsperado, totalObtenido)
	assert.Nil(t, errorObtenido)
}

// En este test compruebo que el GetTotalTickets devuelve un error cuando se pasa un
// string vacío, y este error es el
func TestGetTotalTickets_EmptyDestination(t *testing.T) {
	// Act
	_, errorObtenido := GetTotalTickets("")

	// Assert
	assert.True(t, errors.Is(errorObtenido, ErrEmptyDestination))
}

// En este test compruebo que el GetTotalTickets devuelve un valor en 0
// cuando el destino es inexistente entre los pasajes, y no lanza ningún error,
// ya que considero algo que puede ocurrir normalmente
func TestGetTotalTickets_InexistentDestination(t *testing.T) {
	// Arrange
	totalEsperado := 0

	// Act
	totalObtenido, errorObtenido := GetTotalTickets("Marte")

	// Assert
	assert.Equal(t, totalEsperado, totalObtenido)
	assert.Nil(t, errorObtenido)
}

// Testeo el GetCountByPeriod, en el que espero la cantidad de tickets comprados en distintos
// horarios, y el error obtenido es nil, ya que no hubo errores
func TestGetCountByPeriod(t *testing.T) {
	// Arrange
	totalMañanaEsperado := 267
	totalTardeEsperado := 282
	totalNocheEsperado := 191
	totalMadrugadaEsperado := 260

	// Act
	resultadoObtenido, errorObtenido := GetCountByPeriod()

	// Assert
	assert.Equal(t, totalMadrugadaEsperado, resultadoObtenido.Madrugada)
	assert.Equal(t, totalMañanaEsperado, resultadoObtenido.Mañana)
	assert.Equal(t, totalTardeEsperado, resultadoObtenido.Tarde)
	assert.Equal(t, totalNocheEsperado, resultadoObtenido.Noche)

	assert.Nil(t, errorObtenido)
}

// Compruebo que cuando se pasa un destino valido en AverageDestination,
// este devuelve el promedio de los tickets comprados que van a ese destino, sin errores
func TestAverageDestination_ValidDestination(t *testing.T) {
	// Arrange
	resultadoEsperado := 4

	// Act
	resultadoObtenido, errorObtenido := AverageDestination("Russia")

	// Assert
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.Nil(t, errorObtenido)
}

// Testeo que cuando se pasa como parámetro un string vacío, devuelve un error
func TestAverageDestination_EmptyDestination(t *testing.T) {
	// Act
	_, errorObtenido := AverageDestination("")

	// Assert
	assert.True(t, errors.Is(errorObtenido, ErrEmptyDestination))
}

// Compruebo que cuando se pasa como parámetro un string con destino inexistente,
// devuelve 0 pero no hay errores.
func TestAverageDestination_InexistentDestination(t *testing.T) {
	// Arrange
	resultadoEsperado := 0

	// Act
	resultadoObtenido, errorObtenido := AverageDestination("Marte")

	// Assert
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.Nil(t, errorObtenido)
}
