package lemmy

type PersonBlockView struct {
	Person Person `json:"person" validate:"required"`
	Target Person `json:"target" validate:"required"`
}
