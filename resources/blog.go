package resources

func (src *Source) FetchArticle(key string) (interface{}, error) {
	return src.get("blog", "article", key)
}

func (src *Source) FetchArticles(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("blog", "article", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
