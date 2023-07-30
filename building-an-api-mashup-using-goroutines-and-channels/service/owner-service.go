package service

import (
	"fmt"
	"net/http"
)

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type OwnerService interface {
	FetchData()
}

type fetchOwnerDataService struct{}

func NewOwnerService() CarService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url: %s", ownerServiceUrl)

	resp, _ := client.Get(ownerServiceUrl)

	ownerDataChannel <- resp
}
