package service

import (
	"errors"
	"github/andiahmads/unit-test/mock/entity"
	"github/andiahmads/unit-test/mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepositoy
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}

}
