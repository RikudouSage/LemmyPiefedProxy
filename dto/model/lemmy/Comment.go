package lemmy

type Comment struct {
	ActivityPubId string  `json:"ap_id" validate:"required"`
	Content       string  `json:"content" validate:"required"`
	CreatorId     uint    `json:"creator_id" validate:"required"`
	Deleted       bool    `json:"deleted" validate:"required"`
	Distinguished bool    `json:"distinguished" validate:"required"`
	Id            uint    `json:"id" validate:"required"`
	LanguageId    uint    `json:"language_id" validate:"required"`
	Local         bool    `json:"local" validate:"required"`
	Path          string  `json:"path" validate:"required"`
	PostId        uint    `json:"post_id" validate:"required"`
	Published     string  `json:"published" validate:"required"`
	Removed       bool    `json:"removed" validate:"required"`
	Updated       *string `json:"updated,omitempty"`
}
