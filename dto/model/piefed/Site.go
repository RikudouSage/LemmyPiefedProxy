package piefed

type Site struct {
	ActorId          string            `json:"actor_id" validate:"required"`
	AllLanguages     []LanguageView    `json:"all_languages,omitempty"`
	Description      *string           `json:"description,omitempty"`
	EnableDownvotes  *bool             `json:"enable_downvotes,omitempty"`
	Icon             *string           `json:"icon,omitempty"`
	Name             string            `json:"name" validate:"required"`
	RegistrationMode *RegistrationMode `json:"registration_mode,omitempty"`
	Sidebar          *string           `json:"sidebar,omitempty"`
	UserCount        *uint             `json:"user_count,omitempty"`
}
