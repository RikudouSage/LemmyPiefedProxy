package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertPersonAggregates(in piefed.PersonAggregates) lemmy.PersonAggregates {
	return lemmy.PersonAggregates{
		CommentCount: in.CommentCount,
		PersonId:     in.PersonId,
		PostCount:    in.PostCount,
	}
}
