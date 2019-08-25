package logbook

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type History struct {
}

func (req *History) Default(ctx context.Contexer) (int, interface{}) {
	//req.Setup("history", "History", false)

	return http.StatusOK, nil
}
