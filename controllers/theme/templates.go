package theme

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Templates struct {
}

func (c *Templates) Get(ctx context.Requester) (int, interface{}) {
	//c.Setup("templates", "Templates", false)

	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "html")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Templates) Search(ctx context.Requester) (int, interface{}) {
	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "html")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Templates) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "???", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
