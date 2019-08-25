package stock

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Stock struct {
}

func (req *Stock) Default(ctx context.Contexer) (int, interface{}) {
	//req.Setup("stockhome", "Stock Home", false)

	stocks := []string{
		"Parts",
		"Services",
		"Cars",
	}

	return http.StatusOK, stocks
}
