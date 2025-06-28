package piefed

type CreateCommentRequest struct {
	Body       string `json:"body" validate:"required"`
	PostId     uint   `json:"post_id" validate:"required"`
	ParentId   *uint  `json:"parent_id,omitempty"`
	LanguageId *uint  `json:"language_id,omitempty"`
}
