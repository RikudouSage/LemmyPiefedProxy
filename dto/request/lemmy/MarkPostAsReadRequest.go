package lemmy

type MarkPostAsReadRequest struct {
	PostIds []uint `json:"post_ids" validate:"required"`
	Read    bool   `json:"read" validate:"required"`
}
