package resources

func (src *Source) FetchSecurityReport(key string) (interface{}, error) {
	return src.get("secure", "report", key)
}

func (src *Source) FetchSecurityReports(pagesize string) (interface{}, error) {
	return src.get("secure", "report", pagesize)
}

func (src *Source) FetchProfile(key string) (interface{}, error) {
	return src.get("secure", "profile", key)
}

func (src *Source) FetchProfiles(pagesize string) (interface{}, error) {
	return src.get("secure", "profile", pagesize)
}

func (src *Source) FetchUser(key string) (interface{}, error) {
	return src.get("secure", "user", key)
}

func (src *Source) FetchUsers(pagesize string) (interface{}, error) {
	return src.get("secure", "user", pagesize)
}
