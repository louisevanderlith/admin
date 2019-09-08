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

func (c *Messages) Get(ctx context.Requester) (int, interface{}) {
	//c.Setup("messages", "Messages", true)

	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "message", "A10")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) Search(ctx context.Requester) (int, interface{}) {
	//c.Setup("messages", "Messages", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "message", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) View(ctx context.Requester) (int, interface{}) {
	//c.Setup("messageView", "View Message", false)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comms.API", "message", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
