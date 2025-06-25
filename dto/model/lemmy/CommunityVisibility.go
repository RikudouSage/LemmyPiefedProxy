package lemmy

type CommunityVisibility string

const (
	CommunityVisibilityPublic    CommunityVisibility = "Public"
	CommunityVisibilityLocalOnly CommunityVisibility = "LocalOnly"
)
