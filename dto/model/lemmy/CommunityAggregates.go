package lemmy

type CommunityAggregates struct {
	Comments            uint   `json:"comments" required:"true"`
	CommunityId         uint   `json:"community_id" required:"true"`
	Posts               uint   `json:"posts" required:"true"`
	Published           string `json:"published" required:"true"`
	Subscribers         uint   `json:"subscribers" required:"true"`
	SubscribersLocal    uint   `json:"subscribers_local" required:"true"`
	UsersActiveDay      uint   `json:"users_active_day" required:"true"`
	UsersActiveHalfYear uint   `json:"users_active_half_year" required:"true"`
	UsersActiveMonth    uint   `json:"users_active_month" required:"true"`
	UsersActiveWeek     uint   `json:"users_active_week" required:"true"`
}
