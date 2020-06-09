package resources

func (src *Source) FetchStylesheet(key string) (interface{}, error) {
	return src.get("theme", "stylesheet", key)
}

func (src *Source) FetchStylesheets(pagesize string) (interface{}, error) {
	return src.get("theme", "stylesheet", pagesize)
}

func (src *Source) FetchTemplate(key string) (interface{}, error) {
	return src.get("theme", "template", key)
}

func (src *Source) FetchTemplates(pagesize string) (interface{}, error) {
	return src.get("theme", "template", pagesize)
}
