package folio

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Profiles struct {
}

func (c *Profiles) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("profile", "Profiles", true)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Folio.API", "profile", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Profiles) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("profile", "Profiles", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Folio.API", "profile", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Profiles) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("profileEdit", "Edit Profile", true)
	//c.EnableSave()

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Folio.API", "profile", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Profiles) Create(ctx context.Contexer) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
