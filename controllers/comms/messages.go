package comms

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Messages struct {
}

func (c *Messages) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("messages", "Messages", true)

	var result []interface{}
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "message", "all", "A10")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("messages", "Messages", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "message", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("messageView", "View Message", false)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "message", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
