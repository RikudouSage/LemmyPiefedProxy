package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertSortType(in piefed.SortType) lemmy.SortType {
	return lemmy.SortType(in)
}
