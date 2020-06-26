package resources

func (src *Source) FetchStylesheet(key string) (interface{}, error) {
	return src.get("theme", "stylesheet", key)
}

func (src *Source) FetchStylesheets(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("theme", "stylesheet", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchTemplate(key string) (interface{}, error) {
	return src.get("theme", "template", key)
}

func (src *Source) FetchTemplates(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("theme", "template", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
