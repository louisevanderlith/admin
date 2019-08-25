package curity

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	secure "github.com/louisevanderlith/secure/core"
)

type Users struct {
}

func (c *Users) Default(ctx context.Contexer) (int, interface{}) {
	//.Setup("users", "Users", false)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Secure.API", "user", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Users) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("users", "Users", false)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Secure.API", "user", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Users) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("userView", "View User", true)
	//c.EnableSave()

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})

	resultUser := secure.User{}
	code, err := droxolite.DoGET(ctx.GetMyToken(), &resultUser, ctx.GetInstanceID(), "Secure.API", "user", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	result["User"] = resultUser

	resultRouter := make(map[string]struct{})
	code, err = droxolite.DoGET(ctx.GetMyToken(), &resultRouter, ctx.GetInstanceID(), "Router.API", "memory", "apps")

	if err != nil {
		log.Println(err)
		return code, err
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

	return http.StatusOK, result
}
