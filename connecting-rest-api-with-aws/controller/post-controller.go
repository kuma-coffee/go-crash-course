package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kuma-coffee/go-crash-course/connecting-rest-api-with-aws/cache"
	"github.com/kuma-coffee/go-crash-course/connecting-rest-api-with-aws/entity"
	"github.com/kuma-coffee/go-crash-course/connecting-rest-api-with-aws/errors"
	"github.com/kuma-coffee/go-crash-course/connecting-rest-api-with-aws/service"
)

var (
	postService service.PostService
	PostCache   cache.PostCache
)

type controller struct{}

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	PostCache = cache
	return &controller{}
}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	ID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
	}

	// add cache using redis
	stringID := strconv.Itoa(ID)
	var post *entity.Post = PostCache.Get(stringID)
	if post == nil {
		posts, err := postService.FindByID(ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
			return
		}
		PostCache.Set(stringID, posts)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}

}

func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
	}

	err = postService.Validate(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
	}

	result, err := postService.Create(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*controller) DeletePost(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
	}

	err = postService.Delete(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success delete post")

}
