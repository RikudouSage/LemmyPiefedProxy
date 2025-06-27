package lemmy

type GetReportCountRequest struct {
	CommunityId *uint `json:"community_id,omitempty"`
}
