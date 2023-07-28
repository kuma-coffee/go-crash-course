package service

import (
	"testing"

	"github.com/kuma-coffee/go-crash-course/unit-testing-code-by-mocking-with-testify/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1

	post := entity.Post{ID: 1, Title: "A", Text: "B"}
	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock assertion: behavioral
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)

}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.Post{Title: "A", Text: "B"}
	// Setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.Create(&post)

	// Mock assertion: behavioral
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)

}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "1"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

func TestValidateEmptyPostText(t *testing.T) {
	post := entity.Post{ID: 1, Title: "1", Text: ""}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post text is empty", err.Error())
}
