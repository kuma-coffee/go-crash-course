package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/kuma-coffee/go-crash-course/connectiong-rest-api-with-firebase/entity"
	"github.com/kuma-coffee/go-crash-course/connectiong-rest-api-with-firebase/repositorty"
)

var (
	posts []entity.Post
	repo  repositorty.PostRepository = repositorty.NewPostRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Error getting the posts"}`))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Error unmarshalling the request"}`))
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
