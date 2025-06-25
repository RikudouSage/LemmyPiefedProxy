package lemmy

type CommentView struct {
	BannedFromCommunity        bool              `json:"banned_from_community" validate:"required"`
	Comment                    Comment           `json:"comment" validate:"required"`
	Community                  Community         `json:"community" validate:"required"`
	Counts                     CommentAggregates `json:"counts" validate:"required"`
	Creator                    Person            `json:"creator" validate:"required"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community" validate:"required"`
	CreatorBlocked             bool              `json:"creator_blocked" validate:"required"`
	CreatorIsAdmin             bool              `json:"creator_is_admin" validate:"required"`
	CreatorIsModerator         bool              `json:"creator_is_moderator" validate:"required"`
	MyVote                     *int              `json:"my_vote,omitempty"`
	Post                       Post              `json:"post" validate:"required"`
	Saved                      bool              `json:"saved" validate:"required"`
	Subscribed                 SubscribedType    `json:"subscribed" validate:"required"`
}
