package comms

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Messages struct {
}

func (c *Messages) Get(c *gin.Context) {
	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "messages", "A10")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) Search(c *gin.Context) {
	var result []interface{}
	pagesize := c.Param("pagesize")
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "messages", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "messages", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
