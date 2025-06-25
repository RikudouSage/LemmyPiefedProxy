package lemmy

type ListingType string

const (
	ListingTypeAll           ListingType = "All"
	ListingTypeLocal         ListingType = "Local"
	ListingTypeSubscribed    ListingType = "Subscribed"
	ListingTypeModeratorView ListingType = "ModeratorView"
)
