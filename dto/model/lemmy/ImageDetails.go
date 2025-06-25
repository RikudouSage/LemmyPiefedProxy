package lemmy

type ImageDetails struct {
	ContentType string `json:"content_type" validate:"required"`
	Height      uint   `json:"height" validate:"required"`
	Link        string `json:"link" validate:"required"`
	Width       uint   `json:"width" validate:"required"`
}
