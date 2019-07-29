package vin

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type RegionsController struct {
	xontrols.UICtrl
}

func (c *RegionsController) Get() {
	c.Setup("regions", "VIN Regions", true)
	result := []interface{}{}
	pagesize := c.FindParam(":pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "VIN.API", "region", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *RegionsController) GetEdit() {
	c.Setup("regionEdit", "Edit Region", false)
	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "VIN.API", "region", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
