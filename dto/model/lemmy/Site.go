package lemmy

type Site struct {
	ActorId         string  `json:"actor_id" validate:"required"`
	Banner          *string `json:"banner,omitempty"`
	ContentWarning  *string `json:"content_warning,omitempty"`
	Description     *string `json:"description,omitempty"`
	Icon            *string `json:"icon,omitempty"`
	Id              uint    `json:"id" validate:"required"`
	InboxUrl        string  `json:"inbox_url" validate:"required"`
	InstanceId      uint    `json:"instance_id" validate:"required"`
	LastRefreshedAt string  `json:"last_refreshed_at" validate:"required"`
	Name            string  `json:"name" validate:"required"`
	PublicKey       string  `json:"public_key" validate:"required"`
	Published       string  `json:"published" validate:"required"`
	Sidebar         *string `json:"sidebar,omitempty"`
	Updated         *string `json:"updated,omitempty"`
}
