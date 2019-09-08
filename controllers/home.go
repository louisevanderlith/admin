package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

func Index(ctx context.Requester) (int, interface{}) {
	result := make(map[string]string)
	result["Comms.API"] = "/comms/messages/A10"
	result["Router.API"] = "/router/memory"
	result["Folio.API"] = "/folio/profiles/A10"

	return http.StatusOK, result
}
