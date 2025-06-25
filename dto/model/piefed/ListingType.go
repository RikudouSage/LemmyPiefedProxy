package piefed

type ListingType string

const (
	ListingTypeAll           ListingType = "All"
	ListingTypeLocal         ListingType = "Local"
	ListingTypeSubscribed    ListingType = "Subscribed"
	ListingTypePopular       ListingType = "Popular"
	ListingTypeModeratorView ListingType = "ModeratorView"
)
