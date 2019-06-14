package controllers

import (
	"log"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"

	"github.com/louisevanderlith/mango/control"

	"github.com/louisevanderlith/admin/logic"
)

type BlogController struct {
	control.UIController
}

func NewBlogCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *BlogController {
	result := &BlogController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *BlogController) Get() {
	c.Setup("blogs", "Blogs", true)
	c.CreateSideMenu(logic.GetMenu("/blog"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Blog.API", "article", "non", pagesize)

	c.Serve(result, err)
}

func (c *BlogController) GetCreate() {
	c.Setup("blogCreate", "Create Blog", true)
	c.CreateSideMenu(logic.GetMenu("/blog"))
	c.CreateTopMenu(createBlogTopMenu())

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

func (c *BlogController) GetView() {
	c.Setup("blogView", "View Blog", true)
	c.CreateSideMenu(logic.GetMenu("/blog"))
	c.CreateTopMenu(createBlogTopMenu())

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
	result := control.NewMenu("/blog")
	result.AddItem("btnPreview", "#", "Preview", "fa-globe", nil)
	result.AddItem("btnPublish", "#", "Publish", "fa-bolt", nil)

	return result
}
