package cache

import "github.com/kuma-coffee/go-crash-course/connecting-rest-api-with-aws/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
