package converter

import (
	"LemmyPiefedApi/dto/model/ap"
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertInstanceBlockView(in piefed.InstanceBlockView, siteActor ap.Actor) lemmy.InstanceBlockView {
	return lemmy.InstanceBlockView{
		Instance: ConvertInstance(in.Instance),
		Person:   ConvertPerson(in.Person),
		Site:     ConvertSite(in.Site, siteActor),
	}
}
