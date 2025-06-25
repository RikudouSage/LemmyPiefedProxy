package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
)

func ConvertPostAggregates(in piefed.PostAggregates) lemmy.PostAggregates {
	return lemmy.PostAggregates{
		Comments:          in.Comments,
		Downvotes:         in.Downvotes,
		NewestCommentTime: in.NewestCommentTime,
		PostId:            in.PostId,
		Published:         in.Published,
		Score:             in.Score,
		Upvotes:           in.Upvotes,
	}
}
