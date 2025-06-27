package piefed

import (
	"LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
)

func (receiver *Piefed) GetPosts(request *piefed.GetPostsRequest, headers http.Headers) (*piefedResponse.GetPostsResponse, error) {
	return defaultHandler[piefedResponse.GetPostsResponse](
		receiver,
		"/post/list",
		router.HttpMethodGet,
		request,
		headers,
	)
}

func (receiver *Piefed) GetPost(request *piefed.GetPostRequest, headers http.Headers) (*piefedResponse.GetPostResponse, error) {
	return defaultHandler[piefedResponse.GetPostResponse](
		receiver,
		"/post",
		router.HttpMethodGet,
		request,
		headers,
	)
}
