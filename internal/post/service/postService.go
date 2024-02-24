package service

import (
	"github.com/noo-dev/example/internal/post/model"
	"github.com/noo-dev/example/internal/post/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) GetAllPosts() ([]*model.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) GetPostById(id int) (*model.Post, error) {
	return s.repo.GetOne(id)
}

func (s *PostService) CreatePost(body *model.Post) (*model.Post, error) {
	return s.repo.Create(body)
}
