package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistance(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-27.5945, -48.5477)
	distance := NewDistance(coordFrom, coordTo)
	result := distance.Calculate("km")
	assert.Equal(t, 748.2577773401298, result)
}

func TestCalculateDistance_SamePoint(t *testing.T) {
	coordFrom, _ := NewCoord(-22.9129, -43.2003)
	coordTo, _ := NewCoord(-22.9129, -43.2003)
	distance := NewDistance(coordFrom, coordTo)
	result := distance.Calculate("km")
	assert.Equal(t, 0.0, result)
}
