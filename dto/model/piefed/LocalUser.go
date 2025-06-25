package piefed

type LocalUser struct {
	DefaultListingType ListingType `json:"default_listing_type" validate:"required"`
	DefaultSortType    SortType    `json:"default_sort_type" validate:"required"`
	ShowBotAccounts    bool        `json:"show_bot_accounts" validate:"required"`
	ShowNsfw           bool        `json:"show_nsfw" validate:"required"`
	ShowReadPosts      bool        `json:"show_read_posts" validate:"required"`
	ShowScores         bool        `json:"show_scores" validate:"required"`
}
