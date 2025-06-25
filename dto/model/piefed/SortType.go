package piefed

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
)
