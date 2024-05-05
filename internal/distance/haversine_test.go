package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHaversineDistance_CalculateKm(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-27.5945, -48.5477)
	distance := NewHaversineDistance()
	result := distance.Calculate(coordFrom, coordTo, "km")
	assert.Equal(t, 748.2577773401298, result)
}
func TestHaversineDistance_CalculateMi(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-27.5945, -48.5477)
	distance := NewHaversineDistance()
	result := distance.Calculate(coordFrom, coordTo, "mi")
	assert.Equal(t, 464.8570526938053, result)
}

func TestHaversineDistance_CalculateSamePoint(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-22.9129, -43.2003)
	distance := NewHaversineDistance()
	result := distance.Calculate(coordFrom, coordTo, "km")
	assert.Equal(t, 0.0, result)
}
