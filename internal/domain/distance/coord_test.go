package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCoord(t *testing.T) {
	coord, err := NewCoord(-22.9129, -43.2003)
	assert.Nil(t, err)
	assert.Equal(t, -22.9129, coord.Lat)
	assert.Equal(t, -43.2003, coord.Long)
}

func TestCoordInvalid(t *testing.T) {
	coord, err := NewCoord(-91.0, 100.0)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid latitude")
	assert.Nil(t, coord)

	coord, err = NewCoord(89.0, -181.0)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid longitude")
	assert.Nil(t, coord)
}
