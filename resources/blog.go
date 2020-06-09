package resources

func (src *Source) FetchArticle(key string) (interface{}, error) {
	return src.get("blog", "article", key)
}

func (src *Source) FetchArticles(pagesize string) (interface{}, error) {
	return src.get("blog", "article", pagesize)
}
