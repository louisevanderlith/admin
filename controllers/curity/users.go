package curity

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
	secure "github.com/louisevanderlith/secure/core"
)

type Users struct {
	xontrols.UICtrl
}

func (c *Users) Default() {
	c.Setup("users", "Users", false)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Secure.API", "user", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Users) Search() {
	c.Setup("users", "Users", false)

	var result []interface{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Secure.API", "user", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Users) View() {
	c.Setup("userView", "View User", true)
	c.EnableSave()

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
	}

	result := make(map[string]interface{})

	resultUser := secure.User{}
	code, err := droxolite.DoGET(c.GetMyToken(), &resultUser, c.Settings.InstanceID, "Secure.API", "user", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	result["User"] = resultUser

	resultRouter := make(map[string]struct{})
	code, err = droxolite.DoGET(c.GetMyToken(), &resultRouter, c.Settings.InstanceID, "Router.API", "memory", "apps")

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
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

	c.Serve(http.StatusOK, nil, result)
}
