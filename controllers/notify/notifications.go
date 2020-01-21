package logbook

import (
	"net/http"
)

type Notifications struct {
}

func (req *Notifications) Get(c *gin.Context) {
	//req.Setup("notifications", "Notifications", false)

	return http.StatusNotImplemented, nil
}
