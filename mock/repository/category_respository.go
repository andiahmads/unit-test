package repository

import "github/andiahmads/unit-test/mock/entity"

type CategoryRepositoy interface {
	FindById(id string) *entity.Category
}
