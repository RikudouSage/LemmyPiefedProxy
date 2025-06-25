package piefed

import (
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
)

func (receiver *Piefed) Site(headers http.Headers) (*piefedResponse.GetSiteResponse, error) {
	return defaultHandler[piefedResponse.GetSiteResponse](
		receiver,
		"/site",
		router.HttpMethodGet,
		nil,
		headers,
	)
}
