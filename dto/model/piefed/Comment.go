package piefed

type Comment struct {
	Id            uint    `json:"id" validate:"required"`
	CreatorId     uint    `json:"creator_id" validate:"required"`
	PostId        uint    `json:"post_id" validate:"required"`
	Content       string  `json:"content" validate:"required"`
	Removed       bool    `json:"removed" validate:"required"`
	Published     string  `json:"published" validate:"required"`
	Updated       *string `json:"updated,omitempty"`
	Deleted       bool    `json:"deleted" validate:"required"`
	ActivityPubId string  `json:"ap_id" validate:"required"`
	Local         bool    `json:"local" validate:"required"`
	Path          string  `json:"path" validate:"required"`
	Distinguished bool    `json:"distinguished" validate:"required"`
	LanguageId    uint    `json:"language_id" validate:"required"`
}
