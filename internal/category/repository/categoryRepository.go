package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/noo-dev/example/internal/category/model"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(DB *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: DB}
}

func (r *CategoryRepository) GetAll() ([]*model.Category, error) {
	// write db logic here
	return nil, nil
}

func (r *CategoryRepository) GetOne(id int) (*model.Category, error) {
	// write db logic here
	return nil, nil
}

func (r *CategoryRepository) Create(body *model.Category) (*model.Category, error) {
	// write db logic here
	return nil, nil
}

func (r *CategoryRepository) Update(id int, body *model.Category) error {
	// write db logic here
	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	// write db logic here
	return nil
}
