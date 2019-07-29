package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type DefaultController struct {
	xontrols.UICtrl
}

func (c *DefaultController) Get() {
	c.Setup("default", "Admin", false)

	result := make(map[string]string)
	result["Comms.API"] = "/comms"
	result["Router.API"] = "/memory"
	result["Folio.API"] = "/profiles/A10"

	c.Serve(http.StatusOK, nil, result)
}
