package theme

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Templates struct {
}

func (c *Templates) Get(c *gin.Context) {
	//c.Setup("templates", "Templates", false)

	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "html")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Templates) Search(c *gin.Context) {
	var result []interface{}
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Theme.API", "asset", "html")

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Templates) View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

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
