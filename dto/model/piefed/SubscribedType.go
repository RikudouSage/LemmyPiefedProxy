package piefed

type SubscribedType string

const (
	SubscribedTypeSubscribed    SubscribedType = "Subscribed"
	SubscribedTypeNotSubscribed SubscribedType = "NotSubscribed"
	SubscribedTypePending       SubscribedType = "Pending"
)
