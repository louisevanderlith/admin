package logbook

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Notifications struct {
}

func (req *Notifications) Default(ctx context.Contexer) (int, interface{}) {
	//req.Setup("notifications", "Notifications", false)

	return http.StatusNotImplemented, nil
}
