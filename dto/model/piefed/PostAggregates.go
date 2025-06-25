package piefed

type PostAggregates struct {
	PostId            uint   `json:"post_id" validate:"required"`
	Comments          uint   `json:"comments" validate:"required"`
	Score             int    `json:"score" validate:"required"`
	Upvotes           uint   `json:"upvotes" validate:"required"`
	Downvotes         uint   `json:"downvotes" validate:"required"`
	Published         string `json:"published" validate:"required"`
	NewestCommentTime string `json:"newest_comment_time" validate:"required"`
}
