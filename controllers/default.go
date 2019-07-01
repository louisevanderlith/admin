package controllers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type DefaultController struct {
	control.UIController
}

func NewDefaultCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) beego.ControllerInterface {
	result := &DefaultController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *DefaultController) Get() {
	c.Setup("default", "Admin", false)

	result := make(map[string]string)
	result["Comms.API"] = "/comms"
	result["Router.API"] = "/memory"
	result["Folio.API"] = "/profiles/A10"
	//result[]
	c.Serve(result, nil)
}
