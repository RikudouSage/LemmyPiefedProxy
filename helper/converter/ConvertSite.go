package converter

import (
	"LemmyPiefedApi/dto/model/ap"
	"LemmyPiefedApi/dto/model/lemmy"
	"LemmyPiefedApi/dto/model/piefed"
	"time"
)

func ConvertSite(in piefed.Site, actor ap.Actor) lemmy.Site {
	return lemmy.Site{
		ActorId:     in.ActorId,
		Description: in.Description,
		Icon:        in.Icon,
		InboxUrl:    actor.Inbox,
		Name:        in.Name,
		PublicKey:   actor.PublicKey.PublicKeyPem,
		Sidebar:     in.Sidebar,
		Published:   actor.Published,
		Updated:     actor.Updated,
	}
}

func ConvertSiteToView(in piefed.Site, actor ap.Actor) lemmy.SiteView {
	return lemmy.SiteView{
		LocalSite: lemmy.LocalSite{
			ActorNameMaxLength:     20,
			DefaultPostListingMode: lemmy.PostListingModeList,
			DefaultPostListingType: lemmy.ListingTypeSubscribed,
			DefaultSortType:        lemmy.SortTypeHot,
			DefaultTheme:           "browser",
			EnableDownvotes:        in.EnableDownvotes,
			EnableNsfw:             true,
			FederationEnabled:      true,
			Published:              time.Now().Format("2006-01-02T15:04:05Z07:00"),
			RegistrationMode:       ConvertRegistrationMode(in.RegistrationMode),
			SiteSetup:              true,
		},
		LocalSiteRateLimit: lemmy.LocalSiteRateLimit{
			Published: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		},
		Site: ConvertSite(in, actor),
	}
}
