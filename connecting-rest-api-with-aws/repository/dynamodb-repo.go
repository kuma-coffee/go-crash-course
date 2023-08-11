package repositorty

import (
	"math/rand"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/kuma-coffee/go-crash-course/connecting-rest-api-with-aws/entity"
)

type dynamoDBRepo struct {
	tableName string
}

// NewDynamoDBRepository is the constructor function for the repo
func NewDynamoDBRepository() PostRepository {
	return &dynamoDBRepo{
		tableName: "posts",
	}
}

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Creates a new Session by usin .aws/credentials and .aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (repo *dynamoDBRepo) Save(post *entity.Post) (*entity.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Create random ID
	post.ID = rand.Int63n(100)

	// Transform the post to map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		return nil, err
	}

	// Create the Item Input
	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.tableName),
	}

	// Save the Item into DynamoDB
	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		return nil, err
	}

	// Return the post
	return post, nil
}

func (repo *dynamoDBRepo) FindAll() ([]entity.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input params
	params := &dynamodb.ScanInput{
		TableName: aws.String(repo.tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoDBClient.Scan(params)
	if err != nil {
		return nil, err
	}

	// Create a posts array and add all the existing posts
	var posts []entity.Post
	for _, i := range result.Items {
		post := entity.Post{}

		err := dynamodbattribute.UnmarshalMap(i, &post)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	// Return the post
	return posts, nil
}
func (repo *dynamoDBRepo) FindByID(id int) (*entity.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Get the item by ID
	idString := strconv.Itoa(id)
	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(idString),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Map the dynamodb element to the post struct
	post := entity.Post{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
	if err != nil {
		panic(err)
	}

	return &post, nil
}
func (repo *dynamoDBRepo) Delete(id int) error {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Get the item by ID
	idString := strconv.Itoa(id)
	_, err := dynamoDBClient.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(idString),
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
