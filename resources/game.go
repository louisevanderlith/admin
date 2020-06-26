package resources

func (src *Source) FetchHero(key string) (interface{}, error) {
	return src.get("game", "hero", key)
}

func (src *Source) FetchHeroes(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("game", "hero", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
