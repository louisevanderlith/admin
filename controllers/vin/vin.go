package vin

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type VIN struct {
}

func (c *VIN) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("vins", "VIN Numbers", true)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "admin", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *VIN) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("vins", "VIN Numbers", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "admin", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *VIN) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("vinView", "View VIN", false)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "admin", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
