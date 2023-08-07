package repositorty

import (
	"github.com/kuma-coffee/go-crash-course/using-redis-as-a-cache-for-rest-api/entity"
)

type PostRepository interface {
	SaveSqlite(post *entity.Post) (*entity.Post, error)
	FindAllSqlite() ([]entity.Post, error)
	FindByIDSqlite(id int) (*entity.Post, error)
	DeleteSqlite(id int) error
}
