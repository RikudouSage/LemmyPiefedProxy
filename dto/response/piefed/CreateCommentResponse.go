package piefed

import (
	"LemmyPiefedApi/dto/model/piefed"
)

type CreateCommentResponse struct {
	CommentView piefed.CommentView `json:"comment_view" validate:"required"`
}
