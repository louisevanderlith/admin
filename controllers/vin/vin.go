package vin

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type VINController struct {
	control.UIController
}

func NewVINCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) beego.ControllerInterface {
	result := &VINController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *VINController) Get() {
	c.Setup("vins", "VIN Numbers", true)

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "VIN.API", "admin", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *VINController) GetView() {
	c.Setup("vinView", "View VIN", false)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "VIN.API", "admin", key.String())

	c.Serve(result, err)
}
