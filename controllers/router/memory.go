package router

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
)

type Memory struct {
}

func (c *Memory) Get(c *gin.Context) {
	//c.Setup("memory", "Memory", true)

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Router.API", "memory")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
