package piefed

type LanguageView struct {
	Code string `json:"code" validate:"required"`
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
