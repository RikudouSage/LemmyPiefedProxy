package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertSortType(in piefed.SortType) lemmy.SortType {
	return lemmy.SortType(in)
}

func ReverseConvertSortType(in lemmy.SortType) piefed.SortType {
	switch in {
	case lemmy.SortTypeOld:
	case lemmy.SortTypeMostComments:
	case lemmy.SortTypeControversial:
		return piefed.SortTypeActive
	case lemmy.SortTypeTopYear:
	case lemmy.SortTypeTopAll:
	case lemmy.SortTypeTopThreeMonths:
	case lemmy.SortTypeTopSixMonths:
	case lemmy.SortTypeTopNineMonths:
		return piefed.SortTypeTopMonth
	case lemmy.SortTypeNewComments:
		return piefed.SortTypeNew
	}

	return piefed.SortType(in)
}
