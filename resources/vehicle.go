package resources

func (src *Source) FetchVehicle(key string) (interface{}, error) {
	return src.get("vehicle", "vehicles", key)
}

func (src *Source) FetchVehicles(pagesize string) (interface{}, error) {
	return src.get("vehicle", "vehicles", pagesize)
}
