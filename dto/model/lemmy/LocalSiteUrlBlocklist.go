package lemmy

type LocalSiteUrlBlocklist struct {
	Id        uint    `json:"id" validate:"required"`
	Published string  `json:"published" validate:"required"`
	Updated   *string `json:"updated,omitempty"`
	Url       string  `json:"url" validate:"required"`
}
