package request

type SubRedditRequest struct {
	Name        string `json:"name" validate:"required,min=10,50"`
	Description string `json:"description" validate:"required,min=10,max=50"`
}
