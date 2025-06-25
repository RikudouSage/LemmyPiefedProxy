package piefed

type CommunityModeratorView struct {
	Community Community `json:"community" validate:"required"`
	Moderator Person    `json:"moderator" validate:"required"`
}
