package freight

import (
	"errors"
	"math"

	"github.com/virb30/freight-calculator/internal/distance"
)

type Item struct {
	Volume   float64
	Density  float64
	Quantity int32
}

var EmptyItemsError = errors.New("provide at least one item")
var Distance float64 = 1000
var Factor float64 = 100

func CalculateFreight(items []Item, from *distance.Coord, to *distance.Coord) (error, float64) {
	total := 0.0
	if len(items) == 0 {
		return EmptyItemsError, 0.0
	}

	dist := Distance
	if from != nil && to != nil {
		distanceCalculator := distance.NewDistance(from, to)
		dist = distanceCalculator.Calculate("km")
	}
	for _, item := range items {
		itemFreight := item.Volume * dist * (item.Density / Factor)
		total += itemFreight * float64(item.Quantity)
	}
	total = math.Round(total*100) / 100
	return nil, total
}
