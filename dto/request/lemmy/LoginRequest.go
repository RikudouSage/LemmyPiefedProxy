package lemmy

type LoginRequest struct {
	Password        string `json:"password" validate:"required"`
	TotpToken       string `json:"totp_2fa_token,omitempty"`
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
}
