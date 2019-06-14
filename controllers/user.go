package controllers

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
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

func (c *UserController) GetView() {
	c.Setup("userView", "View User", true)
	c.CreateSideMenu(logic.GetMenu("/user"))
	c.EnableSave()

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})

	resultUser := secure.User{}
	code, err := mango.DoGET(c.GetMyToken(), &resultUser, c.GetInstanceID(), "Secure.API", "user", key.String())

	if err != nil {
		log.Printf("code %v error: %s\n", code, err.Error())
		c.Serve(code, err)
		return
	}

	result["User"] = resultUser

	resultRouter := make(map[string]struct{})
	_, err = mango.DoGET(c.GetMyToken(), &resultRouter, c.GetInstanceID(), "Router.API", "memory", "apps")

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	result["Router"] = resultRouter

	resultOpts := make(map[string]struct{})

	for name := range resultRouter {
		nonItem := ""
		for _, role := range resultUser.Roles {
			if role.ApplicationName == name {
				nonItem = name
				break
			}
		}

		if len(nonItem) == 0 {
			resultOpts[name] = struct{}{}
		}
	}

	result["Options"] = resultOpts

	c.Serve(result, nil)
}
