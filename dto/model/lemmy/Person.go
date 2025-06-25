package lemmy

type Person struct {
	ActorId      string `json:"actor_id" validate:"required"`
	Avatar       string `json:"avatar,omitempty"`
	BanExpires   string `json:"ban_expires,omitempty"`
	Banned       bool   `json:"banned" validate:"required"`
	Banner       string `json:"banner,omitempty"`
	Bio          string `json:"bio,omitempty"`
	BotAccount   bool   `json:"bot_account,omitempty"`
	Deleted      bool   `json:"deleted" validate:"required"`
	DisplayName  string `json:"display_name"`
	Id           uint   `json:"id" validate:"required"`
	InstanceId   uint   `json:"instance_id" validate:"required"`
	Local        bool   `json:"local" validate:"required"`
	MatrixUserId string `json:"matrix_user_id,omitempty"`
	Name         string `json:"name" validate:"required"`
	Published    string `json:"published" validate:"required"`
	Updated      string `json:"updated,omitempty"`
}
