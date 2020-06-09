package resources

func (src *Source) FetchStockCar(key string) (interface{}, error) {
	return src.get("stock", "car", key)
}

func (src *Source) FetchStockCars(pagesize string) (interface{}, error) {
	return src.get("stock", "cars", pagesize)
}

func (src *Source) FetchStockService(key string) (interface{}, error) {
	return src.get("stock", "service", key)
}

func (src *Source) FetchStockServices(pagesize string) (interface{}, error) {
	return src.get("stock", "service", pagesize)
}

func (src *Source) FetchStockPart(key string) (interface{}, error) {
	return src.get("stock", "part", key)
}

func (src *Source) FetchStockParts(pagesize string) (interface{}, error) {
	return src.get("stock", "part", pagesize)
}
