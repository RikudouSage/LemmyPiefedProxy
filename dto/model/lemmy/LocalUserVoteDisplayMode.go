package lemmy

type LocalUserVoteDisplayMode struct {
	Downvotes        bool `json:"downvotes" validate:"required"`
	LocalUserId      uint `json:"local_user_id" validate:"required"`
	Score            bool `json:"score" validate:"required"`
	UpvotePercentage bool `json:"upvote_percentage" validate:"required"`
	Upvotes          bool `json:"upvotes" validate:"required"`
}
