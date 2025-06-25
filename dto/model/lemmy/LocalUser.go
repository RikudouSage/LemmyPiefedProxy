package lemmy

type LocalUser struct {
	AcceptedApplication      bool            `json:"accepted_application" validate:"required"`
	Admin                    bool            `json:"admin" validate:"required"`
	AutoExpand               bool            `json:"auto_expand" validate:"required"`
	BlurNsfw                 bool            `json:"blur_nsfw" validate:"required"`
	CollapseBotComments      bool            `json:"collapse_bot_comments" validate:"required"`
	DefaultListingType       ListingType     `json:"default_listing_type" validate:"required"`
	DefaultSortType          SortType        `json:"default_sort_type" validate:"required"`
	Email                    string          `json:"email,omitempty" validate:"email"`
	EmailVerified            bool            `json:"email_verified" validate:"required"`
	EnableAnimatedImages     bool            `json:"enable_animated_images" validate:"required"`
	EnableKeyboardNavigation bool            `json:"enable_keyboard_navigation" validate:"required"`
	Id                       uint            `json:"id" validate:"required"`
	InfiniteScrollEnabled    bool            `json:"infinite_scroll_enabled" validate:"required"`
	InterfaceLanguage        string          `json:"interface_language" validate:"required"`
	LastDonationNotification string          `json:"last_donation_notification" validate:"required"`
	OpenLinksInNewTab        bool            `json:"open_links_in_new_tab" validate:"required"`
	PersonId                 uint            `json:"person_id" validate:"required"`
	PostListingMode          PostListingMode `json:"post_listing_mode" validate:"required"`
	SendNotificationsToEmail bool            `json:"send_notifications_to_email" validate:"required"`
	ShowAvatars              bool            `json:"show_avatars" validate:"required"`
	ShowBotAccounts          bool            `json:"show_bot_accounts" validate:"required"`
	ShowNsfw                 bool            `json:"show_nsfw" validate:"required"`
	ShowReadPosts            bool            `json:"show_read_posts" validate:"required"`
	ShowScores               bool            `json:"show_scores" validate:"required"`
	Theme                    string          `json:"theme" validate:"required"`
	TotpEnabled              bool            `json:"totp_2fa_enabled" validate:"required"`
}
