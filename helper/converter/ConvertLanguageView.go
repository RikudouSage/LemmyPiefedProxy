package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertLanguageView(in piefed.LanguageView) lemmy.Language {
	return lemmy.Language{
		Code: in.Code,
		Id:   in.Id,
		Name: in.Name,
	}
}
