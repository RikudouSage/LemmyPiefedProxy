package piefed

type PersonAggregates struct {
	CommentCount uint `json:"comment_count" validate:"required"`
	PersonId     uint `json:"person_id" validate:"required"`
	PostCount    uint `json:"post_count" validate:"required"`
}
