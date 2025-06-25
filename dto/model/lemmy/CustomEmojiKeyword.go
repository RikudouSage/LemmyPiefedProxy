package lemmy

type CustomEmojiKeyword struct {
	CustomEmojiId uint   `json:"custom_emoji_id" validate:"required"`
	Keyword       string `json:"keyword" validate:"required"`
}
