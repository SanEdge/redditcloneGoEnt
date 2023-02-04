package request

type PostRequest struct {
	SubredditName string `json:"subreddit" validate:"required"`
	PostName      string `json:"postname" validate:"required"`
	Url           string `json:"url" validate:"required"`
	Description   string `json:"description" validate:"required"`
}
