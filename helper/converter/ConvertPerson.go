package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertPerson(in piefed.Person) lemmy.Person {
	return lemmy.Person{
		ActorId:      in.ActorId,
		Avatar:       in.Avatar,
		BanExpires:   "",
		Banned:       in.Banned,
		Banner:       in.Banner,
		Bio:          "",
		BotAccount:   in.Bot,
		Deleted:      in.Deleted,
		DisplayName:  in.Title,
		Id:           in.Id,
		InstanceId:   in.InstanceId,
		Local:        in.Local,
		MatrixUserId: "",
		Name:         in.Username,
		Published:    in.Published,
		Updated:      "",
	}
}
