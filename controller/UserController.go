package controller

import (
	"LemmyPiefedApi/dto/request/lemmy"
	"LemmyPiefedApi/dto/request/piefed"
	lemmyResponse "LemmyPiefedApi/dto/response/lemmy"
	"LemmyPiefedApi/helper"
	"LemmyPiefedApi/http"
	pfService "LemmyPiefedApi/service/piefed"
	goHttp "net/http"
)

type UserController struct {
	piefed *pfService.Piefed
}

func NewUserController(piefed *pfService.Piefed) *UserController {
	return &UserController{
		piefed: piefed,
	}
}

func (receiver *UserController) Login(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequest[lemmy.LoginRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.Login(&piefed.LoginRequest{
		Username: reqDto.UsernameOrEmail,
		Password: reqDto.Password,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.LoginResponse{
			Jwt: resp.Jwt,
		},
	}, nil
}

func (receiver *UserController) Register(request *http.Request) (*http.Response, error) {
	_, err := helper.ParseRequest[lemmy.RegisterRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	return http.NotImplementedResponse(), nil
}

func (receiver *UserController) GetUnreadCount(request *http.Request) (*http.Response, error) {
	resp, err := receiver.piefed.GetUnreadCount(request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetUnreadCountResponse{
			Mentions:        resp.Mentions,
			PrivateMessages: resp.PrivateMessages,
			Replies:         resp.Replies,
		},
	}, nil
}
