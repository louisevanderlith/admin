package quote

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Invoices struct {
}

func (req *Invoices) Get(ctx context.Requester) (int, interface{}) {
	//req.Setup("invoice", "Invoice", false)

	return http.StatusNotImplemented, nil
}
