package piefed

type CommunityAggregates struct {
	CommunityId        uint   `json:"community_id" required:"true"`
	SubscriptionsCount uint   `json:"subscriptions_count" required:"true"`
	PostCount          uint   `json:"post_count" required:"true"`
	PostReplyCount     uint   `json:"post_reply_count" required:"true"`
	Published          string `json:"published" required:"true"`
}
