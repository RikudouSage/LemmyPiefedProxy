package lemmy

type GetReportCountResponse struct {
	CommentReports        uint  `json:"comment_reports" validate:"required"`
	CommunityId           *uint `json:"community_id,omitempty"`
	PostReports           uint  `json:"post_reports" validate:"required"`
	PrivateMessageReports *uint `json:"private_message_reports,omitempty"`
}
