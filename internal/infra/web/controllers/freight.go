package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/virb30/freight-calculator/internal/domain/distance"
	"github.com/virb30/freight-calculator/internal/infra/web"
	"github.com/virb30/freight-calculator/internal/usecase"
)

type FreightController struct {
	WebServer          *web.WebServer
	DistanceCalculator distance.DistanceInterface
}

func NewFreightController(webServer *web.WebServer, distanceCalculator distance.DistanceInterface) {
	controller := &FreightController{
		WebServer:          webServer,
		DistanceCalculator: distanceCalculator,
	}
	webServer.AddHandler("/freight", controller.GetFreight)
}

func (c *FreightController) GetFreight(w http.ResponseWriter, r *http.Request) {
	var input usecase.CalculateFreightInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	calculateFreight := usecase.NewCalculateFreightUsecase(c.DistanceCalculator)
	output, err := calculateFreight.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
