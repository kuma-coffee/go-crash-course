package main

import (
	"fmt"
	"net/http"

	"github.com/kuma-coffee/go-crash-course/implementing-clean-architecture-in-rest-api/controller"
	router "github.com/kuma-coffee/go-crash-course/implementing-clean-architecture-in-rest-api/http"
	"github.com/kuma-coffee/go-crash-course/implementing-clean-architecture-in-rest-api/repositorty"
	"github.com/kuma-coffee/go-crash-course/implementing-clean-architecture-in-rest-api/service"
)

var (
	postRepository repositorty.PostRepository = repositorty.NewFirestoreRepository()
	postService    service.PostService        = service.NewPostService(postRepository)
	postController controller.PostController  = controller.NewPostController(postService)
	httpRouter     router.Router              = router.NewChiRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
