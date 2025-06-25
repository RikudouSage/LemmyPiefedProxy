package piefed

type LocalUserView struct {
	Counts    PersonAggregates `json:"counts" validate:"required"`
	LocalUser LocalUser        `json:"local_user" validate:"required"`
	Person    Person           `json:"person" validate:"required"`
}
