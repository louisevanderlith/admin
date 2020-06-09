package resources

func (src *Source) FetchRegion(key string) (interface{}, error) {
	return src.get("vin", "region", key)
}

func (src *Source) FetchRegions(pagesize string) (interface{}, error) {
	return src.get("vin", "region", pagesize)
}

func (src *Source) FetchVIN(key string) (interface{}, error) {
	return src.get("vin", "admin", key)
}

func (src *Source) FetchVINs(pagesize string) (interface{}, error) {
	return src.get("vin", "admin", pagesize)
}
