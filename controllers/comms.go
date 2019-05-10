package controllers

import (
	"log"

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
	c.Setup("comms", "Messages", false)
	c.CreateSideMenu(logic.GetMenu("/comms"))

	result := []interface{}{}
	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Comms.API", "message", "all", "A10")

	c.Serve(result, err)
}

func (c *CommsController) GetView() {
	c.Setup("commsView", "View Message", false)
	c.CreateSideMenu(logic.GetMenu("/comms"))

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Comms.API", "message", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
