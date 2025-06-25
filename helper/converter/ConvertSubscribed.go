package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertSubscribedType(in piefed.SubscribedType) lemmy.SubscribedType {
	return lemmy.SubscribedType(in)
}
