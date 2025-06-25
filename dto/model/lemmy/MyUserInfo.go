package lemmy

type MyUserInfo struct {
	CommunityBlocks     []CommunityBlockView     `json:"community_blocks" validate:"required"`
	DiscussionLanguages []uint                   `json:"discussion_languages" validate:"required"`
	Follows             []CommunityFollowerView  `json:"follows" validate:"required"`
	InstanceBlocks      []InstanceBlockView      `json:"instance_blocks" validate:"required"`
	LocalUserView       LocalUserView            `json:"local_user_view" validate:"required"`
	Moderates           []CommunityModeratorView `json:"moderates" validate:"required"`
	PersonBlocks        []PersonBlockView        `json:"person_blocks" validate:"required"`
}
