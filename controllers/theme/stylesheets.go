package theme

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Stylesheets struct {
}

func (c *Stylesheets) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("stylesheets", "Stylesheets", false)

	var result []interface{}
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "css")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Stylesheets) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("stylesheets", "Stylesheets", false)

	var result []interface{}
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "css")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Stylesheets) View(ctx context.Contexer) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "???", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
