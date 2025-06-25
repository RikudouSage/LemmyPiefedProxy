package piefed

type Instance struct {
	Id        uint   `json:"id" validate:"required"`
	Domain    string `json:"domain" validate:"required"`
	Published string `json:"published" validate:"required"`
	Updated   string `json:"updated,omitempty"`
	Software  string `json:"software,omitempty"`
	Version   string `json:"version,omitempty"`
}
