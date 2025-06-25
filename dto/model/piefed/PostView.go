package piefed

type PostView struct {
	Post                       Post           `json:"post" required:"true"`
	Creator                    Person         `json:"creator" required:"true"`
	Community                  Community      `json:"community" required:"true"`
	CreatorBannedFromCommunity bool           `json:"creator_banned_from_community" required:"true"`
	BannedFromCommunity        bool           `json:"banned_from_community" required:"true"`
	CreatorIsModerator         bool           `json:"creator_is_moderator" required:"true"`
	CreatorIsAdmin             bool           `json:"creator_is_admin" required:"true"`
	Counts                     PostAggregates `json:"counts" required:"true"`
	Subscribed                 SubscribedType `json:"subscribed" required:"true"`
	Saved                      bool           `json:"saved" required:"true"`
	ActivityAlert              *bool          `json:"activity_alert,omitempty"`
	Read                       bool           `json:"read" required:"true"`
	Hidden                     bool           `json:"hidden" required:"true"`
	CreatorBlocked             *bool          `json:"creator_blocked,omitempty"`
	MyVote                     *int           `json:"my_vote,omitempty"`
	UnreadComments             uint           `json:"unread_comments" required:"true"`
}
