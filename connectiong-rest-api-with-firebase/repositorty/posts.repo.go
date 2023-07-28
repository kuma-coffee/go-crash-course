package repositorty

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/kuma-coffee/go-crash-course/connectiong-rest-api-with-firebase/entity"
	"google.golang.org/api/option"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct {
}

// New Post Repository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "pragmatic-reviews-e8d1d"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("../pragmatic-reviews-e8d1d-firebase-adminsdk-4bxsm-6b0c976176.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore: %v", err)
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("../pragmatic-reviews-e8d1d-firebase-adminsdk-4bxsm-6b0c976176.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore: %v", err)
	}

	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		fmt.Println(err)
		if err != nil {
			break
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}
	return posts, nil
}
