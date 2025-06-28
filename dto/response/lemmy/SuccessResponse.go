package lemmy

type SuccessResponse struct {
	Success bool `json:"success" validate:"required"`
}
