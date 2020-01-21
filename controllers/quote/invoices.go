package quote

import (
	"net/http"
)

type Invoices struct {
}

func (req *Invoices) Get(c *gin.Context) {
	//req.Setup("invoice", "Invoice", false)

	return http.StatusNotImplemented, nil
}
