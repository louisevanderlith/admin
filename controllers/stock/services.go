package stock

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type ServicesController struct {
	control.UIController
}

func NewServicesCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *ServicesController {
	result := &ServicesController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ServicesController) Get() {
	c.Setup("services", "Services", true)
	c.CreateSideMenu(logic.GetMenu("/services"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "service", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *ServicesController) GetEdit() {
	c.Setup("servicesEdit", "Edit services", false)
	c.CreateSideMenu(logic.GetMenu("/services"))
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "service", key.String())

	c.Serve(result, err)
}
