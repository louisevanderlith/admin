package logbook

import (
	"net/http"
)

type History struct {
}

func (req *History) Get(c *gin.Context) {
	//req.Setup("history", "History", false)

	return http.StatusOK, nil
}
