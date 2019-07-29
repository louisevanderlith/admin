package stock

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type PartsController struct {
	xontrols.UICtrl
}

func (c *PartsController) Get() {
	c.Setup("parts", "Parts", true)

	result := []interface{}{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "part", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *PartsController) GetEdit() {
	c.Setup("partsEdit", "Edit parts", false)

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "part", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
