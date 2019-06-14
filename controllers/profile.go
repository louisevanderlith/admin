package controllers

import (
	"log"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"

	"github.com/louisevanderlith/mango/control"

	"github.com/louisevanderlith/admin/logic"
)

type ProfileController struct {
	control.UIController
}

func NewProfileCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *ProfileController {
	result := &ProfileController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ProfileController) Get() {
	c.Setup("profile", "Profiles", true)
	c.CreateSideMenu(logic.GetMenu("/profile"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Folio.API", "profile", "all", pagesize)

	c.Serve(result, err)
}

func (c *ProfileController) GetEdit() {
	c.Setup("profileEdit", "Edit Profile", true)
	c.CreateSideMenu(logic.GetMenu("/profile"))
	c.EnableSave()
	
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Folio.API", "profile", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
