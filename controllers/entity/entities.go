package entity

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Entities struct {
}

func (c *Entities) Get(c *gin.Context) {
	//c.Setup("entity", "Entity", true)

	var result []interface{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Entity.API", "info", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Entities) Search(c *gin.Context) {
	//c.Setup("entity", "Entity", true)

	var result []interface{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Entity.API", "info", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Entities) View(c *gin.Context) {
	//c.Setup("entityEdit", "Edit Entity", true)
	//c.EnableSave()

	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Entity.API", "info", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, err
}
