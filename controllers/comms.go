package controllers

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type CommsController struct {
	control.UIController
}

func NewCommsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *CommsController {
	result := &CommsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *CommsController) Get() {
	c.Setup("comms", "Messages", true)
	c.CreateSideMenu(logic.GetMenu("/comms"))

	result := []interface{}{}
	err := mango.DoGET(&result, c.GetInstanceID(), "Comms.API", "message", "all", "A10")

	c.Serve(result, err)
}

func (c *CommsController) GetView() {
	c.Setup("commsView", "View Message", true)
	c.CreateSideMenu(logic.GetMenu("/comms"))

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	err = mango.DoGET(&result, c.GetInstanceID(), "Comms.API", "message", key.String())

	c.Serve(result, err)
}
