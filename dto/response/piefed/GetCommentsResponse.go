package piefed

import (
	"LemmyPiefedApi/dto/model/piefed"
)

type GetCommentsResponse struct {
	Comments []piefed.CommentView `json:"comments" validate:"required"`
}
