package lemmy

type Community struct {
	ActorId                 string              `json:"actor_id" validate:"required"`
	Banner                  *string             `json:"banner,omitempty"`
	Deleted                 bool                `json:"deleted" validate:"required"`
	Description             *string             `json:"description,omitempty"`
	Hidden                  bool                `json:"hidden" validate:"required"`
	Icon                    *string             `json:"icon,omitempty"`
	Id                      uint                `json:"id" validate:"required"`
	InstanceId              uint                `json:"instance_id" validate:"required"`
	Local                   bool                `json:"local" validate:"required"`
	Name                    string              `json:"name" validate:"required"`
	Nsfw                    bool                `json:"nsfw" validate:"required"`
	PostingRestrictedToMods bool                `json:"posting_restricted_to_mods" validate:"required"`
	Published               string              `json:"published" validate:"required"`
	Removed                 bool                `json:"removed" validate:"required"`
	Title                   string              `json:"title" validate:"required"`
	Updated                 *string             `json:"updated,omitempty"`
	Visibility              CommunityVisibility `json:"visibility" validate:"required"`
}
