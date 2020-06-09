package resources

func (src *Source) FetchCredits(key string) (interface{}, error) {
	return src.get("xchange", "credit", key)
}
