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
	c.Setup("profile")
	c.CreateSideMenu(logic.GetMenu("/profile"))

	resp, err := mango.GETMessage(c.GetInstanceID(), "Folio.API", "profile", "all", "A10")

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if resp.Failed() {
		c.Serve(nil, resp)
		return
	}

	c.Serve(resp.Data, err)
}

func (c *ProfileController) GetEdit() {
	c.Setup("profileEdit")
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	resp, err := mango.GETMessage(c.GetInstanceID(), "Folio.API", "profile", key.String())

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if resp.Failed() {
		c.Serve(nil, resp)
		return
	}

	c.Serve(resp.Data, err)
}
