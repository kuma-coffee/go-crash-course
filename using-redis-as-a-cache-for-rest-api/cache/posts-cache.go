package cache

import "github.com/kuma-coffee/go-crash-course/using-redis-as-a-cache-for-rest-api/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
