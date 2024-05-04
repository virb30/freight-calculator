package freight

import "errors"

type Item struct {
	Volume   float64
	Density  float64
	Quantity int32
}

var EmptyItemsError = errors.New("provide at least one item")
var Distance float64 = 1000
var Factor float64 = 100

func CalculateFreight(items []Item) (error, float64) {
	total := 0.0
	if len(items) == 0 {
		return EmptyItemsError, 0.0
	}
	for _, item := range items {
		itemFreight := item.Volume * Distance * (item.Density / Factor)
		total += itemFreight * float64(item.Quantity)
	}
	return nil, total
}
