package logic

var uploadURL string

/*
func setImageURLs(instanceID string) {
	if uploadURL == "" {
		setUploadURL(instanceID)
	}

	for i := 0; i < len(obj.PortfolioItems); i++ {
		row := &obj.PortfolioItems[i]

		if row.ImageID != 0 {
			row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
		}
	}

	for i := 0; i < len(obj.Headers); i++ {
		row := &obj.Headers[i]

		if row.ImageID != 0 {
			row.ImageURL = uploadURL + strconv.FormatInt(row.ImageID, 10)
		}
	}
}

func setUploadURL(instanceID string) {
	url, err := mango.GetServiceURL(instanceID, "Artifact.API", true)

	if err != nil {
		log.Print("setImageURLs:", err)
	}

	uploadURL = url + "v1/upload/file/"
}
*/
