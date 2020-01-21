package xchange

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Credits struct {
}

func (c *Credits) Get(c *gin.Context) {
	//c.Setup("credits", "Credits", true)

	return http.StatusNotImplemented, nil
}

func (c *Credits) Search(c *gin.Context) {
	//c.Setup("credits", "Credits", true)

	return http.StatusNotImplemented, nil
}

func (c *Credits) View(c *gin.Context) {
	//c.Setup("creditView", "View Credit", false)

	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "XChange.API", "???", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
