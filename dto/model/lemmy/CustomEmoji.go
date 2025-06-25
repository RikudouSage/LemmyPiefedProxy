package lemmy

type CustomEmoji struct {
	AltText     string `json:"alt_text" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Id          uint   `json:"id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	LocalSiteId uint   `json:"local_site_id" validate:"required"`
	Published   string `json:"published" validate:"required"`
	Shortcode   string `json:"shortcode" validate:"required"`
	Updated     string `json:"updated,omitempty"`
}
