package ap

type Actor struct {
	Inbox     string    `json:"inbox" validate:"required"`
	Published string    `json:"published" validate:"required"`
	Updated   string    `json:"updated,omitempty"`
	PublicKey PublicKey `json:"publicKey" validate:"required"`
}
