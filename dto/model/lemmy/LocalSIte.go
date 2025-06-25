package lemmy

type LocalSite struct {
	ActorNameMaxLength         uint             `json:"actor_name_max_length" validate:"required"`
	ApplicationEmailAdmins     bool             `json:"application_email_admins" validate:"required"`
	ApplicationQuestion        string           `json:"application_question,omitempty"`
	CaptchaDifficulty          string           `json:"captcha_difficulty" validate:"required"`
	CaptchaEnabled             bool             `json:"captcha_enabled" validate:"required"`
	CommunityCreationAdminOnly bool             `json:"community_creation_admin_only" validate:"required"`
	DefaultPostListingMode     PostListingMode  `json:"default_post_listing_mode" validate:"required"`
	DefaultPostListingType     ListingType      `json:"default_post_listing_type" validate:"required"`
	DefaultSortType            SortType         `json:"default_sort_type" validate:"required"`
	DefaultTheme               string           `json:"default_theme" validate:"required"`
	EnableDownvotes            bool             `json:"enable_downvotes" validate:"required"`
	EnableNsfw                 bool             `json:"enable_nsfw" validate:"required"`
	FederationEnabled          bool             `json:"federation_enabled" validate:"required"`
	FederationSignedFetch      bool             `json:"federation_signed_fetch" validate:"required"`
	HideModlogModNames         bool             `json:"hide_modlog_mod_names" validate:"required"`
	Id                         uint             `json:"id" validate:"required"`
	LegalInformation           string           `json:"legal_information,omitempty"`
	PrivateInstance            bool             `json:"private_instance" validate:"required"`
	Published                  string           `json:"published" validate:"required"`
	RegistrationMode           RegistrationMode `json:"registration_mode" validate:"required"`
	ReportsEmailAdmin          bool             `json:"reports_email_admins" validate:"required"`
	ReportsEmailVerification   bool             `json:"reports_email_verification" validate:"required"`
	SiteId                     uint             `json:"site_id" validate:"required"`
	SiteSetup                  bool             `json:"site_setup" validate:"required"`
	SlurFilterRegex            string           `json:"slur_filter_regex,omitempty"`
	Updated                    string           `json:"updated,omitempty"`
}
