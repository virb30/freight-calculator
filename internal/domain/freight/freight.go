package freight

import (
	"errors"
	"math"

	"github.com/virb30/freight-calculator/internal/domain/distance"
)

const (
	Factor = 100
)

type Item struct {
	Volume   float64
	Density  float64
	Quantity int32
}

type Freight struct {
	distanceCalculator distance.DistanceInterface
}

var EmptyItemsError = errors.New("provide at least one item")

func NewFreight(distanceCalculator distance.DistanceInterface) *Freight {
	return &Freight{
		distanceCalculator: distanceCalculator,
	}
}

func (f *Freight) Calculate(items []Item, from *distance.Coord, to *distance.Coord) (error, float64) {
	total := 0.0
	if len(items) == 0 {
		return EmptyItemsError, 0.0
	}
	dist := f.distanceCalculator.Calculate(from, to, "km")
	for _, item := range items {
		itemFreight := item.Volume * dist * (item.Density / Factor)
		total += itemFreight * float64(item.Quantity)
	}
	total = math.Round(total*100) / 100
	return nil, total
}
