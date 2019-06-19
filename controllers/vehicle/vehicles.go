package vehicle

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type VehiclesController struct {
	control.UIController
}

func NewVehiclesCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *VehiclesController {
	result := &VehiclesController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *VehiclesController) Get() {
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

func (c *VehiclesController) GetView() {
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
