package piefed

type GetUnreadCountResponse struct {
	Replies         uint `json:"replies" required:"true"`
	Mentions        uint `json:"mentions" required:"true"`
	PrivateMessages uint `json:"private_messages" required:"true"`
	Other           uint `json:"other" required:"true"`
}
