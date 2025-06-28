package lemmy

type CreateCommentRequest struct {
	Content    string `json:"content" validate:"required"`
	LanguageId *uint  `json:"language_id,omitempty"`
	ParentId   *uint  `json:"parent_id,omitempty"`
	PostId     uint   `json:"post_id" validate:"required"`
}
