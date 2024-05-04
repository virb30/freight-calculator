package freight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculteFreight(t *testing.T) {
	items := []Item{
		Item{Volume: 1000, Density: 0.001, Quantity: 3},
		Item{Volume: 125000, Density: 0.00008, Quantity: 1},
		Item{Volume: 30000, Density: 0.0001, Quantity: 1},
	}
	err, result := CalculateFreight(items)
	assert.Nil(t, err)
	assert.Equal(t, 160.0, result)
}

func TestCalculteFreight_Free(t *testing.T) {
	items := []Item{
		Item{Volume: 0, Density: 0, Quantity: 3},
	}
	err, result := CalculateFreight(items)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, result)
}

func TestCalculteFreight_Error(t *testing.T) {
	items := []Item{}
	err, _ := CalculateFreight(items)
	assert.NotNil(t, err)
	assert.Error(t, EmptyItemsError)
}
