package vehicle

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Vehicles struct {
}

func (c *Vehicles) Get(c *gin.Context) {
	//c.Setup("vehicles", "Vehicles", true)

	var result []interface{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Vehicle.API", "vehicle", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Vehicles) Search(c *gin.Context) {
	//c.Setup("vehicles", "Vehicles", true)

	var result []interface{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Vehicle.API", "vehicle", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Vehicles) View(c *gin.Context) {
	//c.Setup("vehicleView", "View Vehicle", false)
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Vehicle.API", "vehicle", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
