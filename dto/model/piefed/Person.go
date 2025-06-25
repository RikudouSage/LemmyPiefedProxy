package piefed

type Person struct {
	ActorId    string  `json:"actor_id" validate:"required"`
	Avatar     *string `json:"avatar,omitempty"`
	Banned     bool    `json:"banned" validate:"required"`
	Banner     *string `json:"banner,omitempty"`
	Bot        bool    `json:"bot" validate:"required"`
	Deleted    bool    `json:"deleted" validate:"required"`
	Id         uint    `json:"id" validate:"required"`
	InstanceId uint    `json:"instance_id" validate:"required"`
	Local      bool    `json:"local" validate:"required"`
	Published  string  `json:"published" validate:"required"`
	Title      *string `json:"title,omitempty"`
	Username   *string `json:"user_name,omitempty"`
}
