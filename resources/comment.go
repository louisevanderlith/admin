package resources

func (src *Source) FetchCommentMessage(key string) (interface{}, error) {
	return src.get("comment", "message", key)
}

func (src *Source) FetchComments(pagesize string) (interface{}, error) {
	return src.get("comment", "message", pagesize)
}
