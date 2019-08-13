package logbook

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type Notifications struct {
	xontrols.UICtrl
}

func (req *Notifications) Default() {
	req.Setup("notifications", "Notifications", false)

	req.Serve(http.StatusNotImplemented, nil, nil)
}
