package stock

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type Stock struct {
	xontrols.UICtrl
}

func (req *Stock) Default() {
	req.Setup("stockhome", "Stock Home", false)

	stocks := []string{
		"Parts",
		"Services",
		"Cars",
	}

	req.Serve(http.StatusOK, nil, stocks)
}
