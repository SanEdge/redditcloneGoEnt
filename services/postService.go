package services

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/interfaces"
)

type PostService = interfaces.IPostService

type postService struct {
	repository interfaces.IPostRepository
}

func NewPostService(repository interfaces.IPostRepository) *postService {
	return &postService{repository: repository}
}

func (s *postService) Create(email string, input request.PostRequest) (*ent.Post, error) {
	res, err := s.repository.Create(email, input)

	return res, err
}

func (s *postService) FindAll() ([]*ent.Post, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *postService) FindById(id int) (*ent.Post, error) {
	res, err := s.repository.FindById(id)

	return res, err
}

func (s *postService) FindPostBySubreddit(id int) ([]*ent.Post, error) {
	res, err := s.repository.FindPostBySubreddit(id)

	return res, err
}

func (s *postService) FindByEmail(email string) (*ent.User, error) {
	res, err := s.repository.FindByEmail(email)

	return res, err
}
