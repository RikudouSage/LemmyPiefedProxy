package lemmy

type PersonView struct {
	Counts  PersonAggregates `json:"counts" validate:"required"`
	IsAdmin bool             `json:"is_admin" validate:"required"`
	Person  Person           `json:"person" validate:"required"`
}
