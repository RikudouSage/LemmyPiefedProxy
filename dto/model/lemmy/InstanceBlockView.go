package lemmy

type InstanceBlockView struct {
	Instance Instance `json:"instance" validate:"required"`
	Person   Person   `json:"person" validate:"required"`
	Site     Site     `json:"site,omitempty"`
}
