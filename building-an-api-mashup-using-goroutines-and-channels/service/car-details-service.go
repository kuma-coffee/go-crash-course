package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kuma-coffee/go-crash-course/building-an-api-mashup-using-goroutines-and-channels/entity"
)

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

type CarDetailService interface {
	GetDetails() entity.CarDetails
}

type service struct{}

func NewCarDetailsService() CarDetailService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {
	go carService.FetchData()
	go ownerService.FetchData()

	car, _ := getCarData()
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.LastName,
	}
}

func getCarData() (entity.Car, error) {
	read := <-carDataChannel

	var car entity.Car

	err := json.NewDecoder(read.Body).Decode(&car)
	if err != nil {
		fmt.Println(err.Error())
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	read := <-ownerDataChannel

	var owner entity.Owner

	err := json.NewDecoder(read.Body).Decode(&owner)
	if err != nil {
		fmt.Println(err.Error())
	}
	return owner, nil
}
