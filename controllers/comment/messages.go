package comment

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/husk"
)

type Messages struct {
}

func (c *Messages) Get(c *gin.Context) {
	//c.Setup("comments", "Comments", true)

	var result []interface{}
	pagesize := "A10"

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comment.API", "message", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) Search(c *gin.Context) {
	var result []interface{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comment.API", "message", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Messages) View(c *gin.Context) {
	//c.Setup("commentView", "View Comment", false)

	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Comment.API", "message", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
