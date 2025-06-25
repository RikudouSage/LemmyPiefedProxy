package lemmy

type SiteAggregates struct {
	Comments            uint `json:"comments" validate:"required"`
	Communities         uint `json:"communities" validate:"required"`
	Posts               uint `json:"posts" validate:"required"`
	SiteId              uint `json:"site_id" validate:"required"`
	Users               uint `json:"users" validate:"required"`
	UsersActiveDay      uint `json:"users_active_day" validate:"required"`
	UsersActiveHalfYear uint `json:"users_active_half_year" validate:"required"`
	UsersActiveMonth    uint `json:"users_active_month" validate:"required"`
	UsersActiveWeek     uint `json:"users_active_week" validate:"required"`
}
