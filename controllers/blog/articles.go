package blog

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"

	"github.com/louisevanderlith/mango/control"
)

type ArticlesController struct {
	control.UIController
}

func NewArticlesCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) logic.PageUI {
	result := &ArticlesController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ArticlesController) Get() {
	c.Setup("articles", "Articles", true)

	c.CreateTopMenu(c.Ctx, getBlogsTopMenu())

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", "non", pagesize)

	c.Serve(result, err)
}

func (c *ArticlesController) GetCreate() {
	c.Setup("articleCreate", "Create Article", true)

	c.CreateTopMenu(c.Ctx, createBlogTopMenu())
	c.EnableSave()

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}

func (c *ArticlesController) GetView() {
	c.Setup("articleView", "View Article", true)

	c.CreateTopMenu(c.Ctx, createBlogTopMenu())
	c.EnableSave()

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}

func createBlogTopMenu() *control.Menu {
	result := control.NewMenu("/article")
	result.AddItemWithID("btnPreview", "#", "Preview", "fa-globe", nil)
	result.AddItemWithID("btnPublish", "#", "Publish", "fa-bolt", nil)

	return result
}

func getBlogsTopMenu() *control.Menu {
	result := control.NewMenu("/article")
	result.AddItemWithID("btnAdd", "#", "Add Article", "fa-globe", nil)

	return result
}
