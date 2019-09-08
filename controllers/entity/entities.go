package entity

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Entities struct {
}

func (c *Entities) Get(ctx context.Requester) (int, interface{}) {
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

func (c *Entities) Search(ctx context.Requester) (int, interface{}) {
	//c.Setup("entity", "Entity", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Entity.API", "info", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Entities) View(ctx context.Requester) (int, interface{}) {
	//c.Setup("entityEdit", "Edit Entity", true)
	//c.EnableSave()

	key, err := husk.ParseKey(ctx.FindParam("key"))

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
