package service

import (
	"context"

	"github.com/virb30/freight-calculator/internal/domain/distance"
	"github.com/virb30/freight-calculator/internal/pb"
	"github.com/virb30/freight-calculator/internal/usecase"
)

type FreightService struct {
	pb.UnimplementedFreightServiceServer
	DistanceCalculator distance.DistanceInterface
}

func NewFreightService(distanceCalculator distance.DistanceInterface) *FreightService {
	return &FreightService{
		DistanceCalculator: distanceCalculator,
	}
}

func (f *FreightService) CalculateFreight(ctx context.Context, in *pb.CalculateFreightRequest) (*pb.CalculateFreightResponse, error) {

	calculateFreight := usecase.NewCalculateFreightUsecase(f.DistanceCalculator)
	var items []usecase.CalculateFreightItemDTO
	for _, item := range in.Items {
		items = append(items, usecase.CalculateFreightItemDTO{
			Volume:   item.Volume,
			Density:  item.Density,
			Quantity: item.Quantity,
		})
	}
	inputDTO := usecase.CalculateFreightInputDTO{
		Items: items,
		From: usecase.CoordDTO{
			Lat:  in.From.Lat,
			Long: in.From.Long,
		},
		To: usecase.CoordDTO{
			Lat:  in.To.Lat,
			Long: in.To.Long,
		},
	}
	output, err := calculateFreight.Execute(inputDTO)
	if err != nil {
		return nil, err
	}
	calculateFreightResponse := &pb.CalculateFreightResponse{
		Total: output.Total,
	}
	return calculateFreightResponse, nil
}
