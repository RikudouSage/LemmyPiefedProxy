package lemmy

type CustomEmojiView struct {
	CustomEmoji CustomEmoji          `json:"custom_emoji" validate:"required"`
	Keywords    []CustomEmojiKeyword `json:"keywords" validate:"required"`
}
