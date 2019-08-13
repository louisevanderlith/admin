package quote

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type Invoices struct {
	xontrols.UICtrl
}

func (req *Invoices) Default() {
	req.Setup("invoice", "Invoice", false)

	req.Serve(http.StatusNotImplemented, nil, nil)
}
