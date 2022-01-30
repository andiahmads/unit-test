package repository

import (
	"github/andiahmads/unit-test/mock/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoyMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoyMock) FindById(id string) *entity.Category {
	argument := repository.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	} else {
		category := argument.Get(0).(entity.Category)
		return &category
	}
}
