package lemmy

type CommunityView struct {
	BannedFromCommunity bool                `json:"banned_from_community" validate:"required"`
	Blocked             bool                `json:"blocked" validate:"required"`
	Community           Community           `json:"community" validate:"required"`
	Counts              CommunityAggregates `json:"counts" validate:"required"`
	Subscribed          SubscribedType      `json:"subscribed" validate:"required"`
}
