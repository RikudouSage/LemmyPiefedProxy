package lemmy

type CommentAggregates struct {
	ChildCount uint   `json:"child_count" validate:"required"`
	CommentId  uint   `json:"comment_id" validate:"required"`
	Downvotes  uint   `json:"downvotes" validate:"required"`
	Published  string `json:"published" validate:"required"`
	Score      int    `json:"score" validate:"required"`
	Upvotes    uint   `json:"upvotes" validate:"required"`
}
