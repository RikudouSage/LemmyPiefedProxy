package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertPerson(in piefed.Person) lemmy.Person {
	return lemmy.Person{
		ActorId:     in.ActorId,
		Avatar:      in.Avatar,
		Banned:      in.Banned,
		Banner:      in.Banner,
		BotAccount:  in.Bot,
		Deleted:     in.Deleted,
		DisplayName: in.Title,
		Id:          in.Id,
		InstanceId:  in.InstanceId,
		Local:       in.Local,
		Name:        in.Username,
		Published:   in.Published,
	}
}
