package interfaces

import (
	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/renaldyhidayatt/redditgoent/ent"
)

type IPostRepository interface {
	Create(email string, input request.PostRequest) (*ent.Post, error)
	FindById(id int) (*ent.Post, error)
	FindAll() ([]*ent.Post, error)
	FindPostBySubreddit(id int) ([]*ent.Post, error)
	FindByEmail(email string) (*ent.Post, error)
}

type IPostService interface {
	Create(email string, input request.PostRequest) (*ent.Post, error)
	FindById(id int) (*ent.Post, error)
	FindAll() ([]*ent.Post, error)
	FindPostBySubreddit(id int) ([]*ent.Post, error)
	FindByEmail(email string) (*ent.Post, error)
}
