package resources

func (src *Source) FetchAccount(key string) (interface{}, error) {
	return src.get("fund", "accounts", key)
}

func (src *Source) FetchAccounts(pagesize string) (interface{}, error) {
	return src.get("fund", "accounts", pagesize)
}
