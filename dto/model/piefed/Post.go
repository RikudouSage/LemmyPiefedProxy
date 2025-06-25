package piefed

type Post struct {
	Id            uint    `json:"id" validate:"required"`
	Title         string  `json:"title" validate:"required"`
	Url           *string `json:"url,omitempty"`
	Body          *string `json:"body,omitempty"`
	CreatorId     uint    `json:"creator_id" validate:"required"` // todo the response does not actually contain this
	CommunityId   uint    `json:"community_id" validate:"required"`
	Removed       bool    `json:"removed" validate:"required"`
	Locked        bool    `json:"locked" validate:"required"`
	Published     string  `json:"published" validate:"required"`
	Updated       *string `json:"updated,omitempty"`
	Deleted       bool    `json:"deleted" validate:"required"`
	Nsfw          bool    `json:"nsfw" validate:"required"`
	ThumbnailUrl  *string `json:"thumbnail_url,omitempty"`
	ActivityPubId string  `json:"ap_id" validate:"required"`
	Local         bool    `json:"local" validate:"required"`
	LanguageId    uint    `json:"language_id" validate:"required"`
	Sticky        bool    `json:"sticky" validate:"required"`
	AltText       *string `json:"alt_text,omitempty"`
}
