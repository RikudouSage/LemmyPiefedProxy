package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertCommunityFollowerView(in piefed.CommunityFollowerView) lemmy.CommunityFollowerView {
	return lemmy.CommunityFollowerView{
		Community: ConvertCommunity(in.Community),
		Follower:  ConvertPerson(in.Follower),
	}
}
