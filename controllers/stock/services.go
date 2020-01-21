package stock

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Services struct {
}

func (c *Services) Get(c *gin.Context) {
	//c.Setup("services", "Services", true)

	var result []interface{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "service", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Services) Search(c *gin.Context) {
	//c.Setup("services", "Services", true)

	var result []interface{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "service", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Services) View(c *gin.Context) {
	//c.Setup("servicesEdit", "Edit services", false)

	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "service", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
