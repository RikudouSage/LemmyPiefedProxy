package piefed

import (
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
	"encoding/json"
	"io"
	goHttp "net/http"
)

func (receiver *Piefed) Site(headers http.Headers) (*piefedResponse.GetSiteResponse, error) {
	resp, err := receiver.sendRequest(
		"/site",
		router.HttpMethodGet,
		nil,
		headers,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != goHttp.StatusOK {
		var errorResponse *piefedResponse.ErrorResponse
		_ = json.Unmarshal(body, &errorResponse)
		errorResponse.StatusCode = resp.StatusCode

		return nil, errorResponse
	}

	var response *piefedResponse.GetSiteResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
