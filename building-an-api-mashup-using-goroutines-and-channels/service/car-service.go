package service

import (
	"fmt"
	"net/http"
)

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type CarService interface {
	FetchData()
}

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url: %s", carServiceUrl)

	resp, _ := client.Get(carServiceUrl)

	carDataChannel <- resp
}
