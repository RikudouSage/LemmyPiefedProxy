package piefed

import (
	piefedRequest "LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	appHttp "LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
	"encoding/json"
	"io"
	"net/http"
)

func (receiver *Piefed) Login(request *piefedRequest.LoginRequest, headers appHttp.Headers) (*piefedResponse.LoginResponse, error) {
	resp, err := receiver.sendRequest(
		"/user/login",
		router.HttpMethodPost,
		request,
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

	if resp.StatusCode != http.StatusOK {
		var errorResponse *piefedResponse.ErrorResponse
		_ = json.Unmarshal(body, &errorResponse)
		errorResponse.StatusCode = resp.StatusCode

		return nil, errorResponse
	}

	var response *piefedResponse.LoginResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
