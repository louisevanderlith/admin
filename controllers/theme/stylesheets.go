package theme

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Stylesheets struct {
}

func (c *Stylesheets) Get(ctx context.Requester) (int, interface{}) {
	//c.Setup("stylesheets", "Stylesheets", false)

	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "css")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Stylesheets) Search(ctx context.Requester) (int, interface{}) {
	//c.Setup("stylesheets", "Stylesheets", false)

	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "css")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Stylesheets) View(ctx context.Requester) (int, interface{}) {
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
