package artifact

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Uploads struct {
}

func (req *Uploads) Get(ctx context.Requester) (int, interface{}) {
	var result []interface{}
	pagesize := "A10"
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Artifact.API", "upload", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (req *Uploads) Search(ctx context.Requester) (int, interface{}) {
	var result []interface{}
	pagesize := ctx.FindParam("pagesize")
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Artifact.API", "upload", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Uploads) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Artifact.API", "upload", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
