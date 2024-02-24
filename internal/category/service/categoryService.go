package service

import (
	"github.com/noo-dev/example/internal/category/model"
	"github.com/noo-dev/example/internal/category/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]*model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetCategoryById(id int) (*model.Category, error) {
	return s.repo.GetOne(id)
}

func (s *CategoryService) CreateCategory(body *model.Category) (*model.Category, error) {
	return s.repo.Create(body)
}
