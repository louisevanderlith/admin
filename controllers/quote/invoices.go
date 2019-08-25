package quote

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Invoices struct {
}

func (req *Invoices) Default(ctx context.Contexer) (int, interface{}) {
	//req.Setup("invoice", "Invoice", false)

	return http.StatusNotImplemented, nil
}
