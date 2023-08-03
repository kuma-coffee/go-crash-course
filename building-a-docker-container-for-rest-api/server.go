package main

import (
	"building-a-docker-container-for-rest-api/controller"
	router "building-a-docker-container-for-rest-api/http"
	repositorty "building-a-docker-container-for-rest-api/repository"
	"building-a-docker-container-for-rest-api/service"
	"database/sql"
	"log"
	"os"

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

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.DELETE("/delete", postController.DeletePost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
