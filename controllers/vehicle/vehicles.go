package vehicle

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Vehicles struct {
}

func (c *Vehicles) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("vehicles", "Vehicles", true)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Vehicle.API", "vehicle", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Vehicles) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("vehicles", "Vehicles", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Vehicle.API", "vehicle", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Vehicles) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("vehicleView", "View Vehicle", false)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Vehicle.API", "vehicle", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
