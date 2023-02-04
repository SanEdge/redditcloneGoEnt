package interfaces

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
)

type ISubRedditRepository interface {
	Create(email string, input request.SubRedditRequest) (*ent.Subreddit, error)
	FindAll() ([]*ent.Subreddit, error)
	FindById(id int) (*ent.Subreddit, error)
}

type ISubRedditService interface {
	Create(email string, input request.SubRedditRequest) (*ent.Subreddit, error)
	FindAll() ([]*ent.Subreddit, error)
	FindById(id int) (*ent.Subreddit, error)
}
