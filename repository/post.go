package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/ent/post"
	"github.com/renaldyhidayatt/redditgoent/ent/subreddit"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
)

type postRepository struct {
	db      *ent.Client
	context context.Context
}

func NewPostRepository(db *ent.Client, context context.Context) *postRepository {
	return &postRepository{db: db, context: context}
}

func (r *postRepository) Create(email string, input request.PostRequest) (*ent.Post, error) {
	subreddit, err := r.db.Subreddit.Query().Where(subreddit.NameEQ(input.SubredditName)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create post: %w", err)
	}

	user, err := r.db.User.Query().Where(user.EmailEQ(email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result user: %w", err)
	}

	post, err := r.db.Post.Create().SetPostname(input.PostName).SetURL(input.Url).SetDescription(input.Description).AddUser(user).AddSubreddit(subreddit).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create post: %w", err)
	}

	return post, nil
}

func (r *postRepository) FindById(id int) (*ent.Post, error) {
	post, err := r.db.Post.Query().Where(post.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result post: %w", err)
	}

	return post, nil
}

func (r *postRepository) FindAll() ([]*ent.Post, error) {
	post, err := r.db.Post.Query().WithUser().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed results post: %w", err)
	}

	return post, nil
}

func (r *postRepository) FindPostBySubreddit(id int) ([]*ent.Post, error) {
	sub, err := r.db.Subreddit.Query().Where(subreddit.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result post: %w", err)
	}

	post, err := r.db.Post.Query().WithSubreddit(func(sq *ent.SubredditQuery) {
		sq.Where(subreddit.NameEQ(sub.Name))
	}).All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result post: %w", err)
	}

	return post, nil
}

func (r *postRepository) FindByEmail(username string) (*ent.User, error) {
	post, err := r.db.Post.Query().QueryUser().Where(user.UsernameEQ(username)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result post: %w", err)
	}

	return post, nil

}
