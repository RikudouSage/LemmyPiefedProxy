package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertCommunityBlockView(in piefed.CommunityBlockView) lemmy.CommunityBlockView {
	return lemmy.CommunityBlockView{
		Community: ConvertCommunity(in.Community),
		Person:    ConvertPerson(in.Person),
	}
}
