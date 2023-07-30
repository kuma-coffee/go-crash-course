package main

import (
	"github.com/kuma-coffee/go-crash-course/building-an-api-mashup-using-goroutines-and-channels/controller"
	router "github.com/kuma-coffee/go-crash-course/building-an-api-mashup-using-goroutines-and-channels/http"
	"github.com/kuma-coffee/go-crash-course/building-an-api-mashup-using-goroutines-and-channels/service"
)

var (
	carDetailService    service.CarDetailService       = service.NewCarDetailsService()
	carDetailController controller.CarDetailController = controller.NewCarDetailController(carDetailService)
	httpRouter          router.Router                  = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/carDetails", carDetailController.GetCarDetails)
	httpRouter.SERVE(port)
}
