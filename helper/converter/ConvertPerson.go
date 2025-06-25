package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
	"LemmyPiefedApi/helper"
)

func ConvertPerson(in piefed.Person) lemmy.Person {
	return lemmy.Person{
		ActorId:      in.ActorId,
		Avatar:       in.Avatar,
		BanExpires:   nil,
		Banned:       in.Banned,
		Banner:       in.Banner,
		Bio:          nil,
		BotAccount:   &in.Bot,
		Deleted:      in.Deleted,
		DisplayName:  helper.DefaultOnNil(in.Title, ""),
		Id:           in.Id,
		InstanceId:   in.InstanceId,
		Local:        in.Local,
		MatrixUserId: nil,
		Name:         helper.DefaultOnNil(in.Username, ""),
		Published:    in.Published,
		Updated:      nil,
	}
}
