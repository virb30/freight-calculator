package main

import (
	"fmt"

	"github.com/virb30/freight-calculator/configs"
	"github.com/virb30/freight-calculator/internal/domain/distance"
	"github.com/virb30/freight-calculator/internal/infra/grpc"
	"github.com/virb30/freight-calculator/internal/infra/grpc/commands"
	"github.com/virb30/freight-calculator/internal/infra/web"
	"github.com/virb30/freight-calculator/internal/infra/web/controllers"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	distanceCalculator := distance.NewHaversineDistance()
	webserver := web.NewWebServer(config.WebServerPort)
	controllers.NewFreightController(webserver, distanceCalculator)
	fmt.Println("Starting web server on port", config.WebServerPort)
	ch := make(chan error)
	go webserver.Start(ch)

	grpcServer := grpc.NewGrpcServer(config.GrpcServerPort)
	commands.RegisterFreightService(grpcServer, distanceCalculator)
	fmt.Println("Starting grpc server on port", config.GrpcServerPort)
	go grpcServer.Start(ch)
	if err := <-ch; err != nil {
		panic(err)
	}
}
