package main

import (
	"context"
	"fmt"
	"log"
	"time"

	posts_service "github.com/kuma-coffee/grpc-poc/services"
	"google.golang.org/grpc"
)

var (
	serverURL = "localhost:10000"
)

func getGRPCClient() *grpc.ClientConn {
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial(serverURL, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}

	return conn
}

func main() {
	conn := getGRPCClient()

	defer conn.Close()

	client := posts_service.NewPostServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	posts, err := client.GetPost(ctx, &posts_service.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts.GetPosts() {
		fmt.Println(post.Id)
		fmt.Println(post.Title)
		fmt.Println(post.Text)
	}

}
