package piefed

type GetCommentRequest struct {
	Id uint `json:"id" validate:"required"`
}
