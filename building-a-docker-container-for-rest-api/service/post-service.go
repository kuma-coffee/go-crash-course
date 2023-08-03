package service

import (
	"building-a-docker-container-for-rest-api/entity"
	repositorty "building-a-docker-container-for-rest-api/repository"
	"errors"
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
	Delete(id int) error
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
	return postRepo.SaveSqlite(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return postRepo.FindAllSqlite()
}

func (*service) Delete(id int) error {
	return postRepo.DeleteSqlite(id)
}
