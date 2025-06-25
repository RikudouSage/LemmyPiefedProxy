package converter

import (
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
	"time"
)

func ConvertLocalUser(in piefed.LocalUser, person piefed.Person) lemmy.LocalUser {
	return lemmy.LocalUser{
		AcceptedApplication:      true,
		DefaultListingType:       ConvertListingType(in.DefaultListingType),
		DefaultSortType:          ConvertSortType(in.DefaultSortType),
		EmailVerified:            true,
		EnableAnimatedImages:     true,
		EnableKeyboardNavigation: true,
		InfiniteScrollEnabled:    true,
		InterfaceLanguage:        "en",
		LastDonationNotification: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		OpenLinksInNewTab:        true,
		PostListingMode:          lemmy.PostListingModeList,
		ShowAvatars:              true,
		ShowBotAccounts:          in.ShowBotAccounts,
		ShowNsfw:                 in.ShowNsfw,
		ShowReadPosts:            in.ShowReadPosts,
		ShowScores:               in.ShowScores,
		Theme:                    "browser",
		TotpEnabled:              false,
		PersonId:                 person.Id,
	}
}
