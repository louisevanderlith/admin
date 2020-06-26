package resources

func (src *Source) FetchAccount(key string) (interface{}, error) {
	return src.get("fund", "accounts", key)
}

func (src *Source) FetchAccounts(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("fund", "accounts", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
