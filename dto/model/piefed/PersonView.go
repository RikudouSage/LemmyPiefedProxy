package piefed

type PersonView struct {
	Person        Person           `json:"person" validate:"required"`
	Counts        PersonAggregates `json:"counts" validate:"required"`
	IsAdmin       bool             `json:"is_admin" validate:"required"`
	ActivityAlert bool             `json:"activity_alert" validate:"required"`
}
