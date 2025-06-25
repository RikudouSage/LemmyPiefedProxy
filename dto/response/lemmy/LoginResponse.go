package lemmy

type LoginResponse struct {
	Jwt                 string `json:"jwt"`
	RegistrationCreated bool   `json:"registration_created"`
	VerifyEmailSent     bool   `json:"verify_email_sent"`
}
