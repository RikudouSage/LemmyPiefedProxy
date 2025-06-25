package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertRegistrationMode(in piefed.RegistrationMode) lemmy.RegistrationMode {
	return lemmy.RegistrationMode(in)
}
