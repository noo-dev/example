package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/noo-dev/example/internal/post/model"
)

type PostRepository struct {
	DB *sqlx.DB
}

func NewPostRepository(DB *sqlx.DB) *PostRepository {
	return &PostRepository{DB: DB}
}

func (r *PostRepository) GetAll() ([]*model.Post, error) {
	// write db logic here
	return nil, nil
}

func (r *PostRepository) GetOne(id int) (*model.Post, error) {
	// write db logic here
	return nil, nil
}

func (r *PostRepository) Create(body *model.Post) (*model.Post, error) {
	// write db logic here
	return nil, nil
}

func (r *PostRepository) Update(id int, body *model.Post) error {
	// write db logic here
	return nil
}

func (r *PostRepository) Delete(id int) error {
	// write db logic here
	return nil
}
