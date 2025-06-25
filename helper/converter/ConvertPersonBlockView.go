package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertPersonBlockView(in piefed.PersonBlockView) lemmy.PersonBlockView {
	return lemmy.PersonBlockView{
		Person: ConvertPerson(in.Person),
		Target: ConvertPerson(in.Target),
	}
}
