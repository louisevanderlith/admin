package resources

func (src *Source) FetchContent(key string) (interface{}, error) {
	return src.get("cms", "content", key)
}

func (src *Source) FetchCMS(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("cms", "content", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
