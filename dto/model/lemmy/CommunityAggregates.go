package lemmy

type CommunityAggregates struct {
	Comments            uint   `json:"comments" validate:"required"`
	CommunityId         uint   `json:"community_id" validate:"required"`
	Posts               uint   `json:"posts" validate:"required"`
	Published           string `json:"published" validate:"required"`
	Subscribers         uint   `json:"subscribers" validate:"required"`
	SubscribersLocal    uint   `json:"subscribers_local" validate:"required"`
	UsersActiveDay      uint   `json:"users_active_day" validate:"required"`
	UsersActiveHalfYear uint   `json:"users_active_half_year" validate:"required"`
	UsersActiveMonth    uint   `json:"users_active_month" validate:"required"`
	UsersActiveWeek     uint   `json:"users_active_week" validate:"required"`
}
