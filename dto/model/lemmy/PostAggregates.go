package lemmy

type PostAggregates struct {
	Comments          uint   `json:"comments" validate:"required"`
	Downvotes         uint   `json:"downvotes" validate:"required"`
	NewestCommentTime string `json:"newest_comment_time" validate:"required"`
	PostId            uint   `json:"post_id" validate:"required"`
	Published         string `json:"published" validate:"required"`
	Score             int    `json:"score" validate:"required"`
	Upvotes           uint   `json:"upvotes" validate:"required"`
}
