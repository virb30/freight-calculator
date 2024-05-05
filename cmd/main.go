package main

import (
	"fmt"

	"github.com/virb30/freight-calculator/configs"
	"github.com/virb30/freight-calculator/internal/domain/distance"
	"github.com/virb30/freight-calculator/internal/infra/web"
	"github.com/virb30/freight-calculator/internal/infra/web/controllers"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	webserver := web.NewWebServer(config.WebServerPort)
	distanceCalculator := distance.NewHaversineDistance()
	controllers.NewFreightController(webserver, distanceCalculator)
	fmt.Println("Starting web server on port ", config.WebServerPort)
	webserver.Start()
}
