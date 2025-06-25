package lemmy

type LocalUserView struct {
	Counts                   PersonAggregates         `json:"counts" validate:"required"`
	LocalUser                LocalUser                `json:"local_user" validate:"required"`
	LocalUserVoteDisplayMode LocalUserVoteDisplayMode `json:"local_user_vote_display_mode" validate:"required"`
	Person                   Person                   `json:"person" validate:"required"`
}
