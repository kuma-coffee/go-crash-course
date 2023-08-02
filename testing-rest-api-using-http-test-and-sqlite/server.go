package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/controller"
	router "github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/http"
	repositorty "github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/repository"
	"github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Println(err)
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, title VARCHAR(64), text VARCHAR(64))")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table users!")
	}
	statement.Exec()

	var (
		postRepository repositorty.PostRepository = repositorty.NewSqliteRepository(db)
		postService    service.PostService        = service.NewPostService(postRepository)
		postController controller.PostController  = controller.NewPostController(postService)
		httpRouter     router.Router              = router.NewChiRouter()
	)

	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.DELETE("/delete", postController.DeletePost)

	httpRouter.SERVE(port)
}
