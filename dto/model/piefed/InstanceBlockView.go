package piefed

type InstanceBlockView struct {
	Person   Person   `json:"person" validate:"required"`
	Instance Instance `json:"instance" validate:"required"`
	Site     Site     `json:"site,omitempty"`
}
