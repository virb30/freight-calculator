package freight

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/virb30/freight-calculator/internal/domain/distance"
)

func TestCalculteFreight_ConstantDistance(t *testing.T) {
	items := []Item{
		Item{Volume: 1000, Density: 0.001, Quantity: 3},
		Item{Volume: 125000, Density: 0.00008, Quantity: 1},
		Item{Volume: 30000, Density: 0.0001, Quantity: 1},
	}
	distanceCalculator := distance.NewConstantDistance()
	freight := NewFreight(distanceCalculator)
	err, result := freight.Calculate(items, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 160.0, result)
}

func TestCalculteFreight_Free(t *testing.T) {
	items := []Item{
		Item{Volume: 0, Density: 0, Quantity: 3},
	}
	distanceCalculator := distance.NewConstantDistance()
	freight := NewFreight(distanceCalculator)
	err, result := freight.Calculate(items, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, result)
}

func TestCalculteFreight_Error(t *testing.T) {
	items := []Item{}
	distanceCalculator := distance.NewConstantDistance()
	freight := NewFreight(distanceCalculator)
	err, _ := freight.Calculate(items, nil, nil)
	assert.NotNil(t, err)
	assert.Error(t, EmptyItemsError)
}

func TestCalculteFreight_HaversineDistance(t *testing.T) {
	items := []Item{
		Item{Volume: 1000, Density: 0.001, Quantity: 3},
		Item{Volume: 125000, Density: 0.00008, Quantity: 1},
		Item{Volume: 30000, Density: 0.0001, Quantity: 1},
	}
	from, _ := distance.NewCoord(-22.9129, -43.2003)
	to, _ := distance.NewCoord(-27.5945, -48.5477)
	distanceCalculator := distance.NewHaversineDistance()
	freight := NewFreight(distanceCalculator)
	err, result := freight.Calculate(items, from, to)
	assert.Nil(t, err)
	assert.Equal(t, 119.72, result)
}

func TestCalculteFreight_CosinDistance(t *testing.T) {
	items := []Item{
		Item{Volume: 1000, Density: 0.001, Quantity: 3},
		Item{Volume: 125000, Density: 0.00008, Quantity: 1},
		Item{Volume: 30000, Density: 0.0001, Quantity: 1},
	}
	from, _ := distance.NewCoord(-22.9129, -43.2003)
	to, _ := distance.NewCoord(-27.5945, -48.5477)
	distanceCalculator := distance.NewCosinDistance()
	freight := NewFreight(distanceCalculator)
	err, result := freight.Calculate(items, from, to)
	assert.Nil(t, err)
	assert.Equal(t, 119.72, result)
}
