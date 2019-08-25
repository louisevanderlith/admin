package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Home struct {
}

func (c *Home) AcceptsQuery() map[string]string {
	q := make(map[string]string)
	q["access_token"] = "{access_token}"

	return q
}

func (c *Home) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("default", "Admin", false)

	result := make(map[string]string)
	result["Comms.API"] = "/comms/messages/A10"
	result["Router.API"] = "/router/memory"
	result["Folio.API"] = "/folio/profiles/A10"

	return http.StatusOK, result
}
