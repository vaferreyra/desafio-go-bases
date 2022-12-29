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
	totalEsperado := 15

	// Act
	totalObtenido, errorObtenido := GetTotalTickets("Argentina")

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
