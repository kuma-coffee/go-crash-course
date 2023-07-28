package main

import (
	"fmt"
	"net/http"

	"github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/controller"
	router "github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/http"
	"github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/repositorty"
	"github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/service"
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
