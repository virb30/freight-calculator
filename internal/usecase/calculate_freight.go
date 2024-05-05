package usecase

import (
	"github.com/virb30/freight-calculator/internal/domain/distance"
	"github.com/virb30/freight-calculator/internal/domain/freight"
)

type CalculateFreightUsecase struct {
	distanceCalculator distance.DistanceInterface
}

type CalculateFreightItemDTO struct {
	Volume   float64 `json:"volume"`
	Density  float64 `json:"density"`
	Quantity int32   `json:"quantity"`
}

type CoordDTO struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type CalculateFreightInputDTO struct {
	Items []CalculateFreightItemDTO `json:"items"`
	From  CoordDTO                  `json:"from"`
	To    CoordDTO                  `json:"to"`
}

type CalculateFreightOutputDTO struct {
	Total float64 `json:"total"`
}

func NewCalculateFreightUsecase(distanceCalculator distance.DistanceInterface) *CalculateFreightUsecase {
	return &CalculateFreightUsecase{
		distanceCalculator: distanceCalculator,
	}
}

func (u *CalculateFreightUsecase) Execute(input CalculateFreightInputDTO) (CalculateFreightOutputDTO, error) {
	freightCalculator := freight.NewFreight(u.distanceCalculator)
	from, err := distance.NewCoord(input.From.Lat, input.From.Long)
	if err != nil {
		return CalculateFreightOutputDTO{}, err
	}
	to, err := distance.NewCoord(input.To.Lat, input.To.Long)
	if err != nil {
		return CalculateFreightOutputDTO{}, err
	}
	var items []freight.Item
	for _, item := range input.Items {
		items = append(items, freight.Item{
			Volume:   item.Volume,
			Density:  item.Density,
			Quantity: item.Quantity,
		})
	}

	err, total := freightCalculator.Calculate(items, from, to)
	if err != nil {
		return CalculateFreightOutputDTO{}, err
	}

	return CalculateFreightOutputDTO{
		Total: total,
	}, nil
}
