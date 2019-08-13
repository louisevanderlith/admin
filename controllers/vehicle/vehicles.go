package vehicle

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Vehicles struct {
	xontrols.UICtrl
}

func (c *Vehicles) Default() {
	c.Setup("vehicles", "Vehicles", true)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Vehicle.API", "vehicle", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Vehicles) Search() {
	c.Setup("vehicles", "Vehicles", true)

	var result []interface{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Vehicle.API", "vehicle", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Vehicles) View() {
	c.Setup("vehicleView", "View Vehicle", false)
	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Vehicle.API", "vehicle", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
