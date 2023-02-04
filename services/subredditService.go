package services

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/interfaces"
	"github.com/renaldyhidayatt/redditgoent/repository"
)

type SubRedditService = interfaces.ISubRedditService

type subRedditService struct {
	repository repository.SubRedditRepository
}

func NewSubRedditService(repository repository.SubRedditRepository) *subRedditService {
	return &subRedditService{repository: repository}
}

func (s *subRedditService) Create(email string, input request.SubRedditRequest) (*ent.Subreddit, error) {
	res, err := s.repository.Create(email, input)

	return res, err
}

func (s *subRedditService) FindAll() ([]*ent.Subreddit, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *subRedditService) FindById(id int) (*ent.Subreddit, error) {
	res, err := s.repository.FindById(id)

	return res, err
}
