package resources

func (src *Source) FetchCommsMessage(key string) (interface{}, error) {
	return src.get("comms", "message", key)
}

func (src *Source) FetchComms(pagesize string) (interface{}, error) {
	return src.get("comms", "message", pagesize)
}
