package controllers

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type VehicleController struct {
	control.UIController
}

func NewVehicleCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *VehicleController {
	result := &VehicleController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *VehicleController) Get() {
	c.Setup("vehicles", "Vehicles", true)
	c.CreateSideMenu(logic.GetMenu("/vehicle"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Vehicle.API", "vehicle", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *VehicleController) GetView() {
	c.Setup("vehicleView", "View Vehicle", false)
	c.CreateSideMenu(logic.GetMenu("/vehicle"))
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Vehicle.API", "vehicle", key.String())

	c.Serve(result, err)
}
