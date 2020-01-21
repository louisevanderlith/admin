package curity

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
	secure "github.com/louisevanderlith/secure/core"
)

type Users struct {
}

func (c *Users) Get(c *gin.Context) {
	//.Setup("users", "Users", false)

	var result []interface{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Secure.API", "user", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Users) Search(c *gin.Context) {
	//c.Setup("users", "Users", false)

	var result []interface{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Secure.API", "user", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Users) View(c *gin.Context) {
	//c.Setup("userView", "View User", true)
	//c.EnableSave()

	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})

	resultUser := secure.User{}
	code, err := do.GET(ctx.GetMyToken(), &resultUser, ctx.GetInstanceID(), "Secure.API", "user", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	result["User"] = resultUser

	resultRouter := make(map[string]struct{})
	code, err = do.GET(ctx.GetMyToken(), &resultRouter, ctx.GetInstanceID(), "Router.API", "memory", "apps")

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

func (c *Users) Create(c *gin.Context) {
	result := make(map[string]interface{})

	resultRouter := make(map[string]struct{})
	code, err := do.GET(ctx.GetMyToken(), &resultRouter, ctx.GetInstanceID(), "Router.API", "memory", "apps")

	if err != nil {
		log.Println(err)
		return code, err
	}

	result["Router"] = resultRouter

	resultOpts := make(map[string]struct{})

	for name := range resultRouter {
		resultOpts[name] = struct{}{}
	}

	result["Options"] = resultOpts

	return http.StatusOK, result
}

