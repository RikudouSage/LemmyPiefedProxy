package piefed

type LoginResponse struct {
	Jwt string `json:"jwt" validate:"required"`
}
