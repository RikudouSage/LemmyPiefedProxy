package converter

import (
	"LemmyPiefedApi/dto/model/ap"
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
	"LemmyPiefedApi/helper"
)

func ConvertMyUserInfo(in *piefed.MyUserInfo, siteActor ap.Actor) *lemmy.MyUserInfo {
	if in == nil {
		return nil
	}

	return &lemmy.MyUserInfo{
		CommunityBlocks: helper.MapSlice(in.CommunityBlocks, ConvertCommunityBlockView),
		DiscussionLanguages: helper.MapSlice(in.DiscussionLanguages, func(in piefed.LanguageView) uint {
			return in.Id
		}),
		Follows: helper.MapSlice(in.Follows, ConvertCommunityFollowerView),
		InstanceBlocks: helper.MapSlice(in.InstanceBlocks, func(in piefed.InstanceBlockView) lemmy.InstanceBlockView {
			return ConvertInstanceBlockView(in, siteActor)
		}),
		LocalUserView: ConvertLocalUserView(in.LocalUserView),
		Moderates:     helper.MapSlice(in.Moderates, ConvertCommunityModeratorView),
		PersonBlocks:  helper.MapSlice(in.PersonBlocks, ConvertPersonBlockView),
	}
}
