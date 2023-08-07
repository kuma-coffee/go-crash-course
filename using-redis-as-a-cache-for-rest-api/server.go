package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kuma-coffee/go-crash-course/using-redis-as-a-cache-for-rest-api/cache"
	"github.com/kuma-coffee/go-crash-course/using-redis-as-a-cache-for-rest-api/controller"
	router "github.com/kuma-coffee/go-crash-course/using-redis-as-a-cache-for-rest-api/http"
	repositorty "github.com/kuma-coffee/go-crash-course/using-redis-as-a-cache-for-rest-api/repository"
	"github.com/kuma-coffee/go-crash-course/using-redis-as-a-cache-for-rest-api/service"
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
		PostCache      cache.PostCache            = cache.NewRedishCache("localhost:6379", 1, 10)
		postController controller.PostController  = controller.NewPostController(postService, PostCache)
		httpRouter     router.Router              = router.NewChiRouter()
	)

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.DELETE("/delete/{id}", postController.DeletePost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
