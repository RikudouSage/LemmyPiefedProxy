package lemmy

type CommentSortType string

const (
	CommentSortTypeHot           CommentSortType = "Hot"
	CommentSortTypeTop           CommentSortType = "Top"
	CommentSortTypeNew           CommentSortType = "New"
	CommentSortTypeOld           CommentSortType = "Old"
	CommentSortTypeControversial CommentSortType = "Controversial"
)
