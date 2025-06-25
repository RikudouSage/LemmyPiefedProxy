package lemmy

type GetPostRequest struct {
	CommentId *uint `json:"comment_id"`
	Id        *uint `json:"id"`
}
