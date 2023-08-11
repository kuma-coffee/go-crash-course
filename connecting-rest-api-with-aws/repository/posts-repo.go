package repositorty

import (
	"github.com/kuma-coffee/go-crash-course/connecting-rest-api-with-aws/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id int) (*entity.Post, error)
	Delete(id int) error
}
