package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/entity"
	repositorty "github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/repository"
	"github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/service"
	"github.com/stretchr/testify/assert"
)

const (
	ID    int64  = 123
	TITLE string = "Title 1"
	TEXT  string = "Text 1"
)

var (
	db, _                                     = sql.Open("sqlite3", "./users.db")
	postRepo       repositorty.PostRepository = repositorty.NewSqliteRepository(db)
	postSrv        service.PostService        = service.NewPostService(postRepo)
	postController PostController             = NewPostController(postSrv)
)

func TestGetPosts(t *testing.T) {
	// create a new HTTP POST request
	var jsons = []byte(`{"title": "Title 1", "text":"Text 1"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsons))

	// assign HTTP Handler function (controller AddPost function)
	handler := http.HandlerFunc(postController.AddPost)

	// record HTTP Respones (httptest)
	resp := httptest.NewRecorder()

	// dispatch the HTTP request
	handler.ServeHTTP(resp, req)

	// add assertions on the HTTP status code and the response
	status := resp.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// decode the HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(resp.Body)).Decode(&post)
	// assert HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// clean up
	cleanUp(int(post.ID))
}
func TestAddPost(t *testing.T) {
	// insert new post
	setup()

	// create a new HTTP POST request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// assign HTTP Handler function (controller AddPost function)
	handler := http.HandlerFunc(postController.GetPosts)

	// record HTTP Respones (httptest)
	resp := httptest.NewRecorder()

	// dispatch the HTTP request
	handler.ServeHTTP(resp, req)

	// add assertions on the HTTP status code and the response
	status := resp.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// decode the HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(resp.Body)).Decode(&posts)
	// assert HTTP response
	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// clean up
	cleanUp(int(posts[0].ID))
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepo.SaveSqlite(&post)
}

func cleanUp(id int) {
	postRepo.DeleteSqlite(id)
}
