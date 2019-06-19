package comment

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type CommentController struct {
	control.UIController
}

func NewCommentCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *CommentController {
	result := &CommentController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *CommentController) Get() {
	c.Setup("comments", "Comments", true)
	c.CreateSideMenu(logic.GetMenu("/comment"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Comment.API", "message", "all", pagesize)

	c.Serve(result, err)
}

func (c *CommentController) GetView() {
	c.Setup("commentView", "View Comment", false)
	c.CreateSideMenu(logic.GetMenu("/comment"))
	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Comment.API", "message", key.String())

	c.Serve(result, err)
}
