package vin

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Regions struct {
}

func (c *Regions) Get(c *gin.Context) {
	//c.Setup("regions", "VIN Regions", true)
	var result []interface{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "region", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Regions) Search(c *gin.Context) {
	//c.Setup("regions", "VIN Regions", true)
	var result []interface{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "region", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Regions) View(c *gin.Context) {
	//c.Setup("regionEdit", "Edit Region", false)
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "region", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
