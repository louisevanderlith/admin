package controllers

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type UserController struct {
	control.UIController
}

func NewUserCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *UserController {
	result := &UserController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *UserController) Get() {
	c.Setup("users", "Users", false)
	c.CreateSideMenu(logic.GetMenu("/user"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Secure.API", "user", "all", pagesize)

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
