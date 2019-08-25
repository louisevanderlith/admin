package stock

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Cars struct {
}

func (c *Cars) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("cars", "Cars", true)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Cars) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("cars", "Cars", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Cars) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("carsEdit", "Edit car", false)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
