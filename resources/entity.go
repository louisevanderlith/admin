package resources

func (src *Source) FetchEntity(key string) (interface{}, error) {
	return src.get("entity", "entities", key)
}

func (src *Source) FetchEntities(pagesize string) (interface{}, error) {
	return src.get("entity", "entities", pagesize)
}
