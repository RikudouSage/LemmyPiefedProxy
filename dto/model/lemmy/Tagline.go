package lemmy

type Tagline struct {
	Content     string `json:"content" validate:"required"`
	Id          uint   `json:"id" validate:"required"`
	LocalSiteId uint   `json:"local_site_id" validate:"required"`
	Published   string `json:"published" validate:"required"`
	Updated     string `json:"updated,omitempty"`
}
