package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCosinDistance_CalculateSamePoint(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-22.9129, -43.2003)
	distance := NewCosinDistance()
	result := distance.Calculate(coordFrom, coordTo, "km")
	assert.Equal(t, 0.0, result)
}

func TestCosinDistance_CalculateKm(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-27.5945, -48.5477)
	distance := NewCosinDistance()
	result := distance.Calculate(coordFrom, coordTo, "km")
	assert.Equal(t, 748.2577773401306, result)
}

func TestCosinDistance_CalculateMi(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-27.5945, -48.5477)
	distance := NewCosinDistance()
	result := distance.Calculate(coordFrom, coordTo, "mi")
	assert.Equal(t, 464.8570526938058, result)
}
