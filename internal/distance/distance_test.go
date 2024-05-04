package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistance(t *testing.T) {
	coordFrom := NewCoord(-22.9129, -43.2003)
	coordTo := NewCoord(-27.5945, -48.5477)
	distance := NewDistance(coordFrom, coordTo)
	result := distance.Calculate()
	assert.Equal(t, 748.2217780081631, result)
}
