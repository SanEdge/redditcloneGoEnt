package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/ent/comment"
	"github.com/renaldyhidayatt/redditgoent/ent/post"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
	"github.com/renaldyhidayatt/redditgoent/interfaces"
)

type CommentRepository = interfaces.ICommentRepository

type commentRepository struct {
	db      *ent.Client
	context context.Context
}

func NewCommentRepository(db *ent.Client, context context.Context) *commentRepository {
	return &commentRepository{db: db, context: context}
}

func (r *commentRepository) Create(email string, input request.CommentRequest) (*ent.Comment, error) {
	post, err := r.db.Post.Query().Where(post.IDEQ(input.PostId)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query post: %w", err)
	}

	user, err := r.db.User.Query().Where(user.EmailEQ(email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user: %w", err)
	}

	comment, err := r.db.Comment.Create().SetText(input.Text).AddPost(post).AddUser(user).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query comment: %w", err)
	}

	return comment, nil

}

func (r *commentRepository) FindAllCommentForPost(postId int) ([]*ent.Comment, error) {
	post, err := r.db.Post.Query().Where(post.IDEQ(postId)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query post: %w", err)
	}

	comment, err := r.db.Comment.Query().Where(comment.ID(post.ID)).All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query comment: %w", err)
	}

	return comment, nil
}

func (r *commentRepository) FindAllCommentForUser(email string) ([]*ent.User, error) {
	usr, err := r.db.User.Query().Where(user.EmailEQ(email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query user: %w", err)
	}

	comment, err := r.db.Comment.Query().QueryUser().Where(user.EmailEQ(usr.Email)).All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed query comment: %w", err)
	}

	return comment, nil
}
