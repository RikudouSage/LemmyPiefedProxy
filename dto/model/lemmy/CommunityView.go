package lemmy

type CommunityView struct {
	BannedFromCommunity bool                `json:"banned_from_community" required:"true"`
	Blocked             bool                `json:"blocked" required:"true"`
	Community           Community           `json:"community" required:"true"`
	Counts              CommunityAggregates `json:"counts" required:"true"`
	Subscribed          SubscribedType      `json:"subscribed" required:"true"`
}
