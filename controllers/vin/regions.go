package vin

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type RegionsController struct {
	control.UIController
}

func NewRegionsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) logic.PageUI {
	result := &RegionsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *RegionsController) Get() {
	c.Setup("regions", "VIN Regions", true)
	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "VIN.API", "region", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *RegionsController) GetEdit() {
	c.Setup("regionEdit", "Edit Region", false)
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "VIN.API", "region", key.String())

	c.Serve(result, err)
}
