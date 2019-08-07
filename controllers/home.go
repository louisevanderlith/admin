package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type Home struct {
	xontrols.UICtrl
}

func (c *Home) Default() {
	c.Setup("default", "Admin", false)

	result := make(map[string]string)
	result["Comms.API"] = "/comms/messages/A10"
	result["Router.API"] = "/router/memory"
	result["Folio.API"] = "/folio/profiles/A10"

	err := c.Serve(http.StatusOK, nil, result)

	if err != nil {
		log.Println(err)
	}
}
