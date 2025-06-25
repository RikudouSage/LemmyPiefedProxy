package piefed

type CommunityFollowerView struct {
	Community Community `json:"community" validate:"required"`
	Follower  Person    `json:"follower" validate:"required"`
}
