package service

import (
	"errors"
	"math/rand"

	"github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/entity"
	"github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/repositorty"
)

var (
	postRepo repositorty.PostRepository
)

type service struct {
}

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

func NewPostService(repo repositorty.PostRepository) PostService {
	postRepo = repo
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	if post.Text == "" {
		err := errors.New("the post text is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return postRepo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return postRepo.FindAll()
}
