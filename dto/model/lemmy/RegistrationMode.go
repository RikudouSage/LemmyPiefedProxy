package lemmy

type RegistrationMode string

const (
	RegistrationModeClosed             RegistrationMode = "Closed"
	RegistrationModeRequireApplication RegistrationMode = "RequireApplication"
	RegistrationModeOpen               RegistrationMode = "Open"
)
