package stock

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type PartsController struct {
	control.UIController
}

func NewPartsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) beego.ControllerInterface {
	result := &PartsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *PartsController) Get() {
	c.Setup("parts", "Parts", true)

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "part", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}

func (c *PartsController) GetEdit() {
	c.Setup("partsEdit", "Edit parts", false)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "part", key.String())

	c.Serve(result, err)
}
