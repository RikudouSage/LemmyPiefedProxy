package piefed

type CommunityView struct {
	Community           Community           `json:"community" required:"true"`
	Subscribed          SubscribedType      `json:"subscribed" required:"true"`
	Blocked             bool                `json:"blocked" required:"true"`
	Counts              CommunityAggregates `json:"counts" required:"true"`
	BannedFromCommunity bool                `json:"banned_from_community" required:"true"`
	ActivityAlert       bool                `json:"activity_alert" required:"true"`
}
