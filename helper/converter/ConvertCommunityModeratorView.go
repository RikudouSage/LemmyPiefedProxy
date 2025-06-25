package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertCommunityModeratorView(in piefed.CommunityModeratorView) lemmy.CommunityModeratorView {
	return lemmy.CommunityModeratorView{
		Community: ConvertCommunity(in.Community),
		Moderator: ConvertPerson(in.Moderator),
	}
}
