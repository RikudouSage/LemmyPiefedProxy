package lemmy

type Instance struct {
	Domain    string `json:"domain" validate:"required"`
	Id        uint   `json:"id" validate:"required"`
	Published string `json:"published" validate:"required"`
	Software  string `json:"software,omitempty"`
	Updated   string `json:"updated,omitempty"`
	Version   string `json:"version,omitempty"`
}
