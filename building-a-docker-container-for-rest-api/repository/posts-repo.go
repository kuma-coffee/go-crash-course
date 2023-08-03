package repositorty

import "building-a-docker-container-for-rest-api/entity"

type PostRepository interface {
	SaveSqlite(post *entity.Post) (*entity.Post, error)
	FindAllSqlite() ([]entity.Post, error)
	DeleteSqlite(id int) error
}
