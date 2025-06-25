package piefed

type CommentView struct {
	Comment                    Comment           `json:"comment" validate:"required"`
	Creator                    Person            `json:"creator" validate:"required"`
	Post                       Post              `json:"post" validate:"required"`
	Community                  Community         `json:"community" validate:"required"`
	Counts                     CommentAggregates `json:"counts" validate:"required"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community" validate:"required"`
	BannedFromCommunity        bool              `json:"banned_from_community" validate:"required"`
	CreatorIsModerator         bool              `json:"creator_is_moderator" validate:"required"`
	CreatorIsAdmin             bool              `json:"creator_is_admin" validate:"required"`
	Subscribed                 SubscribedType    `json:"subscribed" validate:"required"`
	Saved                      bool              `json:"saved" validate:"required"`
	ActivityAlert              bool              `json:"activity_alert" validate:"required"`
	CreatorBlocked             bool              `json:"creator_blocked" validate:"required"`
	MyVote                     *int              `json:"my_vote,omitempty"`
}
