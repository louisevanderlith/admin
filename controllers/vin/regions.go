package vin

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Regions struct {
}

func (c *Regions) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("regions", "VIN Regions", true)
	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "region", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Regions) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("regions", "VIN Regions", true)
	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "region", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Regions) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("regionEdit", "Edit Region", false)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "region", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
