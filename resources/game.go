package resources

func (src *Source) FetchHero(key string) (interface{}, error) {
	return src.get("game", "hero", key)
}

func (src *Source) FetchHeroes(pagesize string) (interface{}, error) {
	return src.get("game", "hero", pagesize)
}
