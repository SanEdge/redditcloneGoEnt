package services

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/interfaces"
	"github.com/renaldyhidayatt/redditgoent/repository"
)

type CommentService = interfaces.ICommentService

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) *commentService {
	return &commentService{repository: repository}
}

func (s *commentService) Create(email string, input request.CommentRequest) (*ent.Comment, error) {
	res, err := s.repository.Create(email, input)

	return res, err
}

func (s *commentService) FindAllCommentForPost(postId int) ([]*ent.Comment, error) {
	res, err := s.repository.FindAllCommentForPost(postId)

	return res, err
}

func (s *commentService) FindAllCommentForUser(email string) ([]*ent.User, error) {
	res, err := s.repository.FindAllCommentForUser(email)

	return res, err
}
