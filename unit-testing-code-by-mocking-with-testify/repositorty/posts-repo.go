package repositorty

import "github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
