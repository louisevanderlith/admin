package stock

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Cars struct {
	xontrols.UICtrl
}

func (c *Cars) Default() {
	c.Setup("cars", "Cars", true)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "car", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Cars) Search() {
	c.Setup("cars", "Cars", true)

	var result []interface{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "car", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Cars) View() {
	c.Setup("carsEdit", "Edit car", false)
	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "car", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
