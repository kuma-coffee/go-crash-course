package repositorty

import (
	"github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/entity"
)

type PostRepository interface {
	SaveSqlite(post *entity.Post) (*entity.Post, error)
	FindAllSqlite() ([]entity.Post, error)
	DeleteSqlite(id int) error
}
