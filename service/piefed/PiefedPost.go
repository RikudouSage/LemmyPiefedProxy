package piefed

import (
	"LemmyPiefedApi/dto/request/piefed"
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
	"encoding/json"
	"io"
	goHttp "net/http"
)

func (receiver *Piefed) GetPosts(request *piefed.GetPostsRequest, headers http.Headers) (*piefedResponse.GetPostsResponse, error) {
	resp, err := receiver.sendRequest(
		"/post/list",
		router.HttpMethodGet,
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

	if resp.StatusCode != goHttp.StatusOK {
		var errorResponse *piefedResponse.ErrorResponse
		_ = json.Unmarshal(body, &errorResponse)
		errorResponse.StatusCode = resp.StatusCode

		return nil, errorResponse
	}

	var response *piefedResponse.GetPostsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (receiver *Piefed) GetPost(request *piefed.GetPostRequest, headers http.Headers) (*piefedResponse.GetPostResponse, error) {
	resp, err := receiver.sendRequest(
		"/post",
		router.HttpMethodGet,
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

	if resp.StatusCode != goHttp.StatusOK {
		var errorResponse *piefedResponse.ErrorResponse
		_ = json.Unmarshal(body, &errorResponse)
		errorResponse.StatusCode = resp.StatusCode

		return nil, errorResponse
	}

	var response *piefedResponse.GetPostResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
