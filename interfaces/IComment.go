package interfaces

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
)

type ICommentRepository interface {
	Create(email string, input request.CommentRequest) (*ent.Comment, error)
	FindAllCommentForPost(postId int) ([]*ent.Comment, error)
	FindAllCommentForUser(email string) ([]*ent.User, error)
}

type ICommentService interface {
	Create(email string, input request.CommentRequest) (*ent.Comment, error)
	FindAllCommentForPost(postId int) ([]*ent.Comment, error)
	FindAllCommentForUser(email string) ([]*ent.User, error)
}
