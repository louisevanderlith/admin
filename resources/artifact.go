package resources

func (src *Source) FetchUpload(key string) (interface{}, error) {
	return src.get("artifact", "upload", key)
}

func (src *Source) FetchUploads(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("artifact", "upload", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
