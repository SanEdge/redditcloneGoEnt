package repository

import (
	"context"
	"fmt"

	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/ent/subreddit"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
	"github.com/renaldyhidayatt/redditgoent/interfaces"
)

type SubRedditRepository = interfaces.ISubRedditRepository

type subRedditRepository struct {
	db      *ent.Client
	context context.Context
}

func NewSubRedditRepository(db *ent.Client, context context.Context) *subRedditRepository {
	return &subRedditRepository{db: db, context: context}
}

func (r *subRedditRepository) Create(email string, input request.SubRedditRequest) (*ent.Subreddit, error) {
	_, err := r.db.Subreddit.Query().Where(subreddit.NameEQ(input.Name)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create subreddit: %w", err)
	}

	user, err := r.db.User.Query().Where(user.EmailEQ(email)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed get user: %w", err)
	}

	subreddit, err := r.db.Subreddit.Create().SetName(input.Name).SetDescription(input.Description).AddUser(user).Save(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed create subreddit: %w", err)
	}

	return subreddit, nil

}

func (r *subRedditRepository) FindAll() ([]*ent.Subreddit, error) {

	subreddit, err := r.db.Subreddit.Query().WithPosts().All(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed results subreddit: %w", err)
	}

	return subreddit, nil
}

func (r *subRedditRepository) FindById(id int) (*ent.Subreddit, error) {
	subreddit, err := r.db.Subreddit.Query().Where(subreddit.IDEQ(id)).First(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed result subreddit: %w", err)
	}

	return subreddit, nil
}
