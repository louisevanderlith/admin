package resources

func (src *Source) FetchUpload(key string) (interface{}, error) {
	return src.get("theme", "upload", key)
}

func (src *Source) FetchUploads(pagesize string) (interface{}, error) {
	return src.get("theme", "upload", pagesize)
}
