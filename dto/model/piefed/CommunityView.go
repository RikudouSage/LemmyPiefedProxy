package piefed

type CommunityView struct {
	Community           Community           `json:"community" validate:"required"`
	Subscribed          SubscribedType      `json:"subscribed" validate:"required"`
	Blocked             bool                `json:"blocked" validate:"required"`
	Counts              CommunityAggregates `json:"counts" validate:"required"`
	BannedFromCommunity bool                `json:"banned_from_community" validate:"required"`
	ActivityAlert       bool                `json:"activity_alert" validate:"required"`
}
