package router

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
)

type Memory struct {
}

func (c *Memory) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("memory", "Memory", true)

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Router.API", "memory")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
