package logbook

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type History struct {
	xontrols.UICtrl
}

func (req *History) Default() {
	req.Setup("history", "History", false)

	req.Serve(http.StatusNotImplemented, nil, nil)
}
