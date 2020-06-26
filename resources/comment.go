package resources

func (src *Source) FetchCommentMessage(key string) (interface{}, error) {
	return src.get("comment", "message", key)
}

func (src *Source) FetchComments(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("comment", "message", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
