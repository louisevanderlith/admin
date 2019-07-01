package entity

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"

	"github.com/louisevanderlith/mango/control"
)

type EntitiesController struct {
	control.UIController
}

func NewEntitiesCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) beego.ControllerInterface {
	result := &EntitiesController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *EntitiesController) Get() {
	c.Setup("entity", "Entity", true)

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Entity.API", "info", "all", pagesize)

	c.Serve(result, err)
}

func (c *EntitiesController) GetEdit() {
	c.Setup("entityEdit", "Edit Entity", true)
	c.EnableSave()

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Entity.API", "info", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
