package logbook

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Notifications struct {
}

func (req *Notifications) Get(ctx context.Requester) (int, interface{}) {
	//req.Setup("notifications", "Notifications", false)

	return http.StatusNotImplemented, nil
}
