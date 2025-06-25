package lemmy

type GetUnreadCountResponse struct {
	Mentions        uint `json:"mentions" validate:"required"`
	PrivateMessages uint `json:"private_messages" validate:"required"`
	Replies         uint `json:"replies" validate:"required"`
}
