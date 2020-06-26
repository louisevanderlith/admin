package resources

func (src *Source) FetchCommsMessage(key string) (interface{}, error) {
	return src.get("comms", "message", key)
}

func (src *Source) FetchComms(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("comms", "message", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil

}
