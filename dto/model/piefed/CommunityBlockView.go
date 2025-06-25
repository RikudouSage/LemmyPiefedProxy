package piefed

type CommunityBlockView struct {
	Community Community `json:"community" validate:"required"`
	Person    Person    `json:"person" validate:"required"`
}
