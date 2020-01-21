package stock

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Cars struct {
}

func (c *Cars) Get(c *gin.Context) {
	//c.Setup("cars", "Cars", true)

	var result []interface{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Cars) Search(c *gin.Context) {
	//c.Setup("cars", "Cars", true)

	var result []interface{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Cars) View(c *gin.Context) {
	//c.Setup("carsEdit", "Edit car", false)
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
