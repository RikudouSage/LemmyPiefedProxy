package lemmy

type SortType string

const (
	SortTypeActive         SortType = "Active"
	SortTypeHot            SortType = "Hot"
	SortTypeNew            SortType = "New"
	SortTypeOld            SortType = "Old"
	SortTypeTopDay         SortType = "TopDay"
	SortTypeTopWeek        SortType = "TopWeek"
	SortTypeTopMonth       SortType = "TopMonth"
	SortTypeTopYear        SortType = "TopYear"
	SortTypeTopAll         SortType = "TopAll"
	SortTypeMostComments   SortType = "MostComments"
	SortTypeNewComments    SortType = "NewComments"
	SortTypeTopHour        SortType = "TopHour"
	SortTypeTopSixHour     SortType = "TopSixHour"
	SortTypeTopTwelveHour  SortType = "TopTwelveHour"
	SortTypeTopThreeMonths SortType = "TopThreeMonths"
	SortTypeTopSixMonths   SortType = "TopSixMonths"
	SortTypeTopNineMonths  SortType = "TopNineMonths"
	SortTypeControversial  SortType = "Controversial"
	SortTypeScaled         SortType = "Scaled"
)
