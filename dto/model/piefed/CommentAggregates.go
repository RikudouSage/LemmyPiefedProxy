package piefed

type CommentAggregates struct {
	CommentId  uint   `json:"comment_id" validate:"required"`
	Score      int    `json:"score" validate:"required"`
	Upvotes    uint   `json:"upvotes" validate:"required"`
	Downvotes  uint   `json:"downvotes" validate:"required"`
	Published  string `json:"published" validate:"required"`
	ChildCount uint   `json:"child_count" validate:"required"`
}
