package service

import (
	"github/andiahmads/unit-test/mock/entity"
	"github/andiahmads/unit-test/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepositoy = &repository.CategoryRepositoyMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepositoy}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepositoy.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_Success(t *testing.T) {
	category := entity.Category{
		ID:   "1",
		Name: "Mock Testing",
	}
	categoryRepositoy.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)

}
