package controllers

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango/control"
)

type DefaultController struct {
	control.UIController
}

func NewDefaultCtrl(ctrlMap *control.ControllerMap) *DefaultController {
	result := &DefaultController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *DefaultController) Get() {
	c.Setup("default", "Admin", true)
	c.CreateSideMenu(logic.GetMenu("/"))

	result := make(map[string]string)
	result["Comms.API"] = "/comms"
	result["Router.API"] = "/memory"
	result["Folio.API"] = "/profiles/A10"
	//result[]
	c.Serve(result, nil)
}
