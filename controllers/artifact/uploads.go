package artifact

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Uploads struct {
}

func (req *Uploads) Default(ctx context.Contexer) (int, interface{}) {
	//ctx.Setup("uploads", "Uploads", true)

	var result []interface{}
	pagesize := "A10"
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Artifact.API", "upload", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (req *Uploads) Search(ctx context.Contexer) (int, interface{}) {
	//req.Setup("uploads", "Uploads", true)

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Artifact.API", "upload", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Uploads) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("uploadView", "View Upload", false)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Artifact.API", "upload", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
