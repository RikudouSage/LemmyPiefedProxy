package piefed

type CommunityAggregates struct {
	CommunityId        uint   `json:"community_id" validate:"required"`
	SubscriptionsCount uint   `json:"subscriptions_count" validate:"required"`
	PostCount          uint   `json:"post_count" validate:"required"`
	PostReplyCount     uint   `json:"post_reply_count" validate:"required"`
	Published          string `json:"published" validate:"required"`
}
