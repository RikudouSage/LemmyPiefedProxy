package controller

import (
	lemmyModel "LemmyPiefedApi/dto/model/lemmy"
	piefedModel "LemmyPiefedApi/dto/model/piefed"
	"LemmyPiefedApi/dto/request/lemmy"
	"LemmyPiefedApi/dto/request/piefed"
	lemmyResponse "LemmyPiefedApi/dto/response/lemmy"
	"LemmyPiefedApi/helper"
	"LemmyPiefedApi/helper/converter"
	"LemmyPiefedApi/http"
	pfService "LemmyPiefedApi/service/piefed"
	goHttp "net/http"
)

type PostController struct {
	piefed *pfService.Piefed
}

func NewPostController(piefed *pfService.Piefed) *PostController {
	return &PostController{
		piefed: piefed,
	}
}

func (receiver *PostController) GetPosts(request *http.Request) (*http.Response, error) {
	reqDto, err := helper.ParseRequestQuery[lemmy.GetPostsRequest](request)
	if err != nil {
		return helper.ConvertValidationErrorsToResponse(err), nil
	}

	resp, err := receiver.piefed.GetPosts(&piefed.GetPostsRequest{
		Type: helper.SafeDereference(reqDto.Type, func(in lemmyModel.ListingType) *piefedModel.ListingType {
			return helper.ToPointer(converter.ReverseConvertListingType(in))
		}),
		Sort: helper.SafeDereference(reqDto.Sort, func(in lemmyModel.SortType) *piefedModel.SortType {
			return helper.ToPointer(converter.ReverseConvertSortType(in))
		}),
		Page:          reqDto.Page,
		Limit:         reqDto.Limit,
		CommunityId:   reqDto.CommunityId,
		PersonId:      nil,
		CommunityName: reqDto.CommunityName,
		LikedOnly:     reqDto.LikedOnly,
	}, request.Headers)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: &lemmyResponse.GetPostsResponse{
			NextPage: resp.NextPage,
			Posts:    helper.MapSlice(resp.Posts, converter.ConvertPostView),
		},
	}, nil
}
