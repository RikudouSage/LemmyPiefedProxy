package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertCommunityView(in piefed.CommunityView) lemmy.CommunityView {
	return lemmy.CommunityView{
		BannedFromCommunity: in.BannedFromCommunity,
		Blocked:             in.Blocked,
		Community:           ConvertCommunity(in.Community),
		Counts:              ConvertCommunityAggregates(in.Counts),
		Subscribed:          ConvertSubscribedType(in.Subscribed),
	}
}
