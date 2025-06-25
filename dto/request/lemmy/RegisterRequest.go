package lemmy

type RegisterRequest struct {
	ApplicationQuestionAnswer string `json:"answer,omitempty"`
	CaptchaAnswer             string `json:"captcha_answer,omitempty"`
	CaptchaUuid               string `json:"captcha_uuid,omitempty"`
	Email                     string `json:"email,omitempty"`
	Honeypot                  string `json:"honeypot,omitempty"`
	Password                  string `json:"password" validate:"required"`
	PasswordVerify            string `json:"password_verify" validate:"required"`
	ShowNsfw                  bool   `json:"show_nsfw"`
	Username                  string `json:"username" validate:"required"`
}
