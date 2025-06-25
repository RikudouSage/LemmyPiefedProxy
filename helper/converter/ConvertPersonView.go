package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertPersonView(in piefed.PersonView) lemmy.PersonView {
	return lemmy.PersonView{
		Counts:  ConvertPersonAggregates(in.Counts),
		IsAdmin: in.IsAdmin,
		Person:  ConvertPerson(in.Person),
	}
}
