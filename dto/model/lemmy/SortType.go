package lemmy

type SortType string

const (
	SortTypeActive        SortType = "Active"
	SortTypeHot           SortType = "Hot"
	SortTypeNew           SortType = "New"
	SortTypeTopHour       SortType = "TopHour"
	SortTypeTopSixHour    SortType = "TopSixHour"
	SortTypeTopTwelveHour SortType = "TopTwelveHour"
	SortTypeTopDay        SortType = "TopDay"
	SortTypeTopWeek       SortType = "TopWeek"
	SortTypeTopMonth      SortType = "TopMonth"
	SortTypeScaled        SortType = "Scaled"

	SortTypeOld            SortType = "Old"
	SortTypeTopYear        SortType = "TopYear"
	SortTypeTopAll         SortType = "TopAll"
	SortTypeMostComments   SortType = "MostComments"
	SortTypeNewComments    SortType = "NewComments"
	SortTypeTopThreeMonths SortType = "TopThreeMonths"
	SortTypeTopSixMonths   SortType = "TopSixMonths"
	SortTypeTopNineMonths  SortType = "TopNineMonths"
	SortTypeControversial  SortType = "Controversial"
)
