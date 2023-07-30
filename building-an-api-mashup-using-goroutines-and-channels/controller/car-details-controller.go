package controller

import (
	"encoding/json"
	"net/http"

	"github.com/kuma-coffee/go-crash-course/building-an-api-mashup-using-goroutines-and-channels/service"
)

type controller struct{}

var (
	carDetailService service.CarDetailService
)

type CarDetailController interface {
	GetCarDetails(w http.ResponseWriter, r *http.Request)
}

func NewCarDetailController(service service.CarDetailService) CarDetailController {
	carDetailService = service
	return &controller{}
}

func (*controller) GetCarDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := carDetailService.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
