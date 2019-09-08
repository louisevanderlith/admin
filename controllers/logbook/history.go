package logbook

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type History struct {
}

func (req *History) Get(ctx context.Requester) (int, interface{}) {
	//req.Setup("history", "History", false)

	return http.StatusOK, nil
}
