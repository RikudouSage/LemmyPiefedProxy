package piefed

import (
	piefedRequest "LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	appHttp "LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
)

func (receiver *Piefed) Login(request *piefedRequest.LoginRequest, headers appHttp.Headers) (*piefedResponse.LoginResponse, error) {
	return defaultHandler[piefedResponse.LoginResponse](
		receiver,
		"/user/login",
		router.HttpMethodPost,
		request,
		headers,
	)
}

func (receiver *Piefed) GetUnreadCount(headers appHttp.Headers) (*piefedResponse.GetUnreadCountResponse, error) {
	return defaultHandler[piefedResponse.GetUnreadCountResponse](
		receiver,
		"/user/unread_count",
		router.HttpMethodGet,
		nil,
		headers,
	)
}
