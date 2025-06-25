package piefed

type Post struct {
	Id            uint    `json:"id" required:"true"`
	Title         string  `json:"title" required:"true"`
	Url           *string `json:"url,omitempty"`
	Body          *string `json:"body,omitempty"`
	CreatorId     uint    `json:"creator_id" required:"true"` // todo the response does not actually contain this
	CommunityId   uint    `json:"community_id" required:"true"`
	Removed       bool    `json:"removed" required:"true"`
	Locked        bool    `json:"locked" required:"true"`
	Published     string  `json:"published" required:"true"`
	Updated       *string `json:"updated,omitempty"`
	Deleted       bool    `json:"deleted" required:"true"`
	Nsfw          bool    `json:"nsfw" required:"true"`
	ThumbnailUrl  *string `json:"thumbnail_url,omitempty"`
	ActivityPubId string  `json:"ap_id" required:"true"`
	Local         bool    `json:"local" required:"true"`
	LanguageId    uint    `json:"language_id" required:"true"`
	Sticky        bool    `json:"sticky" required:"true"`
	AltText       *string `json:"alt_text,omitempty"`
}
