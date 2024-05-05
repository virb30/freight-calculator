package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/virb30/freight-calculator/internal/domain/distance"
	"github.com/virb30/freight-calculator/internal/domain/freight"
)

func TestCalculateFreight(t *testing.T) {
	input := CalculateFreightInputDTO{
		Items: []CalculateFreightItemDTO{
			CalculateFreightItemDTO{Volume: 1000, Density: 0.001, Quantity: 3},
			CalculateFreightItemDTO{Volume: 125000, Density: 0.00008, Quantity: 1},
			CalculateFreightItemDTO{Volume: 30000, Density: 0.0001, Quantity: 1},
		},
		From: CoordDTO{
			Lat:  -22.9129,
			Long: -43.2003,
		},
		To: CoordDTO{
			Lat:  -27.5945,
			Long: -48.5477,
		},
	}
	usecase := NewCalculateFreightUsecase(distance.NewConstantDistance())
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, 160.0, output.Total)
}

func TestCalculateFreight_InvalidLatitudeCoordinate(t *testing.T) {
	input := CalculateFreightInputDTO{
		Items: []CalculateFreightItemDTO{
			CalculateFreightItemDTO{Volume: 1000, Density: 0.001, Quantity: 3},
			CalculateFreightItemDTO{Volume: 125000, Density: 0.00008, Quantity: 1},
			CalculateFreightItemDTO{Volume: 30000, Density: 0.0001, Quantity: 1},
		},
		From: CoordDTO{
			Lat:  -100,
			Long: -43.2003,
		},
		To: CoordDTO{
			Lat:  -27.5945,
			Long: -48.5477,
		},
	}
	usecase := NewCalculateFreightUsecase(distance.NewConstantDistance())
	_, err := usecase.Execute(input)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid latitude")
}

func TestCalculateFreight_InvalidLongitudeCoordinate(t *testing.T) {
	input := CalculateFreightInputDTO{
		Items: []CalculateFreightItemDTO{
			CalculateFreightItemDTO{Volume: 1000, Density: 0.001, Quantity: 3},
			CalculateFreightItemDTO{Volume: 125000, Density: 0.00008, Quantity: 1},
			CalculateFreightItemDTO{Volume: 30000, Density: 0.0001, Quantity: 1},
		},
		From: CoordDTO{
			Lat:  -22.9129,
			Long: -200,
		},
		To: CoordDTO{
			Lat:  -27.5945,
			Long: -48.5477,
		},
	}
	usecase := NewCalculateFreightUsecase(distance.NewConstantDistance())
	_, err := usecase.Execute(input)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid longitude")
}

func TestCalculateFreight_InvalidItems(t *testing.T) {
	input := CalculateFreightInputDTO{
		Items: []CalculateFreightItemDTO{},
		From: CoordDTO{
			Lat:  -22.9129,
			Long: -43.2003,
		},
		To: CoordDTO{
			Lat:  -27.5945,
			Long: -48.5477,
		},
	}
	usecase := NewCalculateFreightUsecase(distance.NewConstantDistance())
	_, err := usecase.Execute(input)
	assert.Error(t, err)
	assert.ErrorContains(t, err, freight.EmptyItemsError.Error())
}

func TestCalculateFreight_Free(t *testing.T) {
	input := CalculateFreightInputDTO{
		Items: []CalculateFreightItemDTO{
			CalculateFreightItemDTO{Volume: 1000, Density: 0.001, Quantity: 3},
			CalculateFreightItemDTO{Volume: 125000, Density: 0.00008, Quantity: 1},
			CalculateFreightItemDTO{Volume: 30000, Density: 0.0001, Quantity: 1},
		},
		From: CoordDTO{},
		To:   CoordDTO{},
	}
	usecase := NewCalculateFreightUsecase(distance.NewHaversineDistance())
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, output.Total)
}
