package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	// "fmt"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	return c.filebasedDb.StoreCategory(*Category)
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	return c.filebasedDb.UpdateCategory(id, category)
}

func (c *categoryRepository) Delete(id int) error {
	return c.filebasedDb.DeleteCategory(id)
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	return c.filebasedDb.GetCategoryByID(id)
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	return c.filebasedDb.GetCategories()
}