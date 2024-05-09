package commands

import (
	"github.com/virb30/freight-calculator/internal/domain/distance"
	"github.com/virb30/freight-calculator/internal/infra/grpc"
	"github.com/virb30/freight-calculator/internal/pb"
	"github.com/virb30/freight-calculator/internal/service"
)

func RegisterFreightService(grpcServer *grpc.GrpcServer, distanceCalculator distance.DistanceInterface) {
	freightService := service.NewFreightService(distanceCalculator)
	pb.RegisterFreightServiceServer(grpcServer.GrpcServer, freightService)
}
