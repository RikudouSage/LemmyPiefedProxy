package lemmy

type GetCommentRequest struct {
	Id uint `json:"id" validate:"required"`
}
