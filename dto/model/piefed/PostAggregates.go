package piefed

type PostAggregates struct {
	PostId            uint   `json:"post_id" required:"true"`
	Comments          uint   `json:"comments" required:"true"`
	Score             int    `json:"score" required:"true"`
	Upvotes           uint   `json:"upvotes" required:"true"`
	Downvotes         uint   `json:"downvotes" required:"true"`
	Published         string `json:"published" required:"true"`
	NewestCommentTime string `json:"newest_comment_time" required:"true"`
}
