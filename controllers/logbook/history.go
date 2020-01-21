package logbook

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type History struct {
}

func (req *History) Get(c *gin.Context) {
	//req.Setup("history", "History", false)

	return http.StatusOK, nil
}
