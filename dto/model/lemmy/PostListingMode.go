package lemmy

type PostListingMode string

const (
	PostListingModeList      PostListingMode = "List"
	PostListingModeCard      PostListingMode = "Card"
	PostListingModeSmallCard PostListingMode = "SmallCard"
)
