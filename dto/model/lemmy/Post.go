package lemmy

type Post struct {
	AltText           *string `json:"alt_text,omitempty"`
	ActivityPubId     string  `json:"ap_id" validate:"required"`
	Body              *string `json:"body,omitempty"`
	CommunityId       uint    `json:"community_id" validate:"required"`
	CreatorId         uint    `json:"creator_id" validate:"required"`
	Deleted           bool    `json:"deleted" validate:"required"`
	EmbedDescription  *string `json:"embed_description,omitempty"`
	EmbedTitle        *string `json:"embed_title,omitempty"`
	EmbedVideoUrl     *string `json:"embed_video_url,omitempty"`
	FeaturedCommunity bool    `json:"featured_community" validate:"required"`
	FeaturedLocal     bool    `json:"featured_local" validate:"required"`
	Id                uint    `json:"id" validate:"required"`
	LanguageId        uint    `json:"language_id" validate:"required"`
	Local             bool    `json:"local" validate:"required"`
	Locked            bool    `json:"locked" validate:"required"`
	Name              string  `json:"name" validate:"required"`
	Nsfw              bool    `json:"nsfw" validate:"required"`
	Published         string  `json:"published" validate:"required"`
	Removed           bool    `json:"removed" validate:"required"`
	ThumbnailUrl      *string `json:"thumbnail_url,omitempty"`
	Updated           *string `json:"updated,omitempty"`
	Url               *string `json:"url,omitempty"`
	UrlContentType    *string `json:"url_content_type,omitempty"`
}
