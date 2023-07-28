package repositorty

import "github.com/kuma-coffee/go-crash-course/implementing-clean-architecture-in-rest-api/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
