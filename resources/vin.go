package resources

func (src *Source) FetchRegion(key string) (map[string]interface{}, error) {
	res, err := src.get("vin", "regions", key)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchRegions(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("vin", "regions", pagesize)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchVIN(key string) (map[string]interface{}, error) {
	res, err := src.get("vin", "admin", key)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchVINs(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("vin", "admin", pagesize)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(map[string]interface{}), nil
}
