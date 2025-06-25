package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertListingType(in piefed.ListingType) lemmy.ListingType {
	return lemmy.ListingType(in)
}
