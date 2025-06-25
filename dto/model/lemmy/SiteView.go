package lemmy

type SiteView struct {
	Counts             SiteAggregates     `json:"counts" validate:"required"`
	LocalSite          LocalSite          `json:"local_site" validate:"required"`
	LocalSiteRateLimit LocalSiteRateLimit `json:"local_site_rate_limit" validate:"required"`
	Site               Site               `json:"site" validate:"required"`
}
