package resources

func (src *Source) FetchVehicle(key string) (interface{}, error) {
	return src.get("vehicle", "vehicles", key)
}

func (src *Source) FetchVehicles(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("vehicle", "vehicles", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
