package repositorty

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	firebase "firebase.google.com/go"
// 	"github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/entity"
// 	"google.golang.org/api/option"
// )

// type repo struct {
// }

// // New Firestore Repository
// func NewFirestoreRepository() PostRepository {
// 	return &repo{}
// }

// const (
// 	collectionName string = "posts"
// )

// func (*repo) Save(post *entity.Post) (*entity.Post, error) {
// 	ctx := context.Background()
// 	opt := option.WithCredentialsFile("../pragmatic-reviews.json")
// 	app, err := firebase.NewApp(ctx, nil, opt)
// 	if err != nil {
// 		return nil, fmt.Errorf("error initializing app: %v", err)
// 	}

// 	client, err := app.Firestore(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error initializing firestore: %v", err)
// 	}

// 	defer client.Close()

// 	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
// 		"ID":    post.ID,
// 		"Title": post.Title,
// 		"Text":  post.Text,
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed adding a new post: %v", err)
// 		return nil, err
// 	}

// 	return post, nil
// }

// func (*repo) FindAll() ([]entity.Post, error) {
// 	ctx := context.Background()
// 	opt := option.WithCredentialsFile("../pragmatic-reviews.json")
// 	app, err := firebase.NewApp(ctx, nil, opt)
// 	if err != nil {
// 		return nil, fmt.Errorf("error initializing app: %v", err)
// 	}

// 	client, err := app.Firestore(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error initializing firestore: %v", err)
// 	}

// 	defer client.Close()

// 	var posts []entity.Post
// 	iterator := client.Collection(collectionName).Documents(ctx)
// 	for {
// 		doc, err := iterator.Next()
// 		if err != nil {
// 			break
// 		}

// 		post := entity.Post{
// 			ID:    doc.Data()["ID"].(int64),
// 			Title: doc.Data()["Title"].(string),
// 			Text:  doc.Data()["Text"].(string),
// 		}

// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }
