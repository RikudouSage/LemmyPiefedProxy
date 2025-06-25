package lemmy

import "LemmyPiefedApi/dto/model/lemmy"

type GetPostsResponse struct {
	NextPage *string          `json:"next_page,omitempty"`
	Posts    []lemmy.PostView `json:"posts" validate:"required"`
}
