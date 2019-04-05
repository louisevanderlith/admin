package controllers

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"

	"github.com/louisevanderlith/mango/control"

	"github.com/louisevanderlith/admin/logic"
)

type ProfileController struct {
	control.UIController
}

func NewProfileCtrl(ctrlMap *control.ControllerMap) *ProfileController {
	result := &ProfileController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ProfileController) Get() {
	c.Setup("profile", "Profiles", true)
	c.CreateSideMenu(logic.GetMenu("/profile"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")
	fail, err := mango.DoGET(&result, c.GetInstanceID(), "Folio.API", "profile", "all", pagesize)

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if fail != nil {
		c.Serve(nil, fail)
		return
	}

	c.Serve(result, nil)
}

func (c *ProfileController) GetEdit() {
	c.Setup("profileEdit", "Edit Profile", true)
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	fail, err := mango.DoGET(&result, c.GetInstanceID(), "Folio.API", "profile", key.String())

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if fail != nil {
		c.Serve(nil, fail)
		return
	}

	c.Serve(result, nil)
}
