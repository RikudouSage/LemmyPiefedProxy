package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertCommunityAggregates(in piefed.CommunityAggregates) lemmy.CommunityAggregates {
	return lemmy.CommunityAggregates{
		Comments:            in.PostReplyCount,
		CommunityId:         in.CommunityId,
		Posts:               in.PostCount,
		Published:           in.Published,
		Subscribers:         in.SubscriptionsCount,
		SubscribersLocal:    0,
		UsersActiveDay:      0,
		UsersActiveHalfYear: 0,
		UsersActiveMonth:    0,
		UsersActiveWeek:     0,
	}
}
