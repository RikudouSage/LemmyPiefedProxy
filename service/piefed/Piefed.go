package piefed

import (
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/helper"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
	"encoding/json"
	"fmt"
	"io"
	goHttp "net/http"
)

const applicationJson = "application/json"

func defaultHandler[TResponse any](
	piefed *Piefed,
	urlPath string,
	httpMethod router.HttpMethod,
	request any,
	headers http.Headers,
) (*TResponse, error) {
	resp, err := piefed.sendRequest(
		urlPath,
		httpMethod,
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

	var response TResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type Piefed struct {
	instance string
}

func NewPiefed(instance string) *Piefed {
	return &Piefed{
		instance: instance,
	}
}

func (receiver *Piefed) url() string {
	return fmt.Sprintf("https://%s/api/alpha", receiver.instance)
}

func (receiver *Piefed) sendRequest(
	path string,
	method router.HttpMethod,
	request any,
	headers http.Headers,
) (*goHttp.Response, error) {
	var body io.Reader
	queryString := ""
	if request != nil {
		if method == router.HttpMethodGet {
			marshalled, err := helper.MarshalToQueryString(request)
			if err != nil {
				return nil, err
			}
			queryString = "?" + marshalled
		} else {
			body = helper.MarshalToReader(request)
		}
	}

	req, err := goHttp.NewRequest(
		string(method),
		receiver.url()+path+queryString,
		body,
	)
	if err != nil {
		return nil, err
	}

	req.Header = helper.MapMap(headers, func(value, key string) []string {
		return []string{value}
	})
	req.Header.Set("Content-Type", applicationJson)
	req.Header.Del("Content-Length")

	return goHttp.DefaultClient.Do(req)
}
