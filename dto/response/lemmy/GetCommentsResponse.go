package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

type GetCommentsResponse struct {
	Comments []lemmy.CommentView `json:"comments" required:"true"`
}
