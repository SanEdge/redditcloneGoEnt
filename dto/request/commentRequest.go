package request

type CommentRequest struct {
	PostId int    `json:"post_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}
