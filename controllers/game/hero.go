package game

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type HeroController struct {
	control.UIController
}

func NewHeroCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) beego.ControllerInterface {
	result := &HeroController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *HeroController) Get() {
	c.Setup("heroes", "Heroes", true)

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Game.API", "hero", "all", pagesize)

	c.Serve(result, err)
}

func (c *HeroController) GetEdit() {
	c.Setup("heroView", "View Hero", true)
	c.EnableSave()

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Game.API", "hero", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
