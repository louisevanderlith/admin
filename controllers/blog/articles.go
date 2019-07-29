package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type ArticlesController struct {
	xontrols.UICtrl
}

func (c *ArticlesController) Get() {
	c.Setup("articles", "Articles", true)

	c.CreateTopMenu(getBlogsTopMenu())

	result := []interface{}{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", "non", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *ArticlesController) GetEdit() {
	c.Setup("articleCreate", "Edit Article", true)

	c.CreateTopMenu(createBlogTopMenu())
	c.EnableSave()

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *ArticlesController) GetView() {
	c.Setup("articleView", "View Article", true)

	c.CreateTopMenu(createBlogTopMenu())
	c.EnableSave()

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func createBlogTopMenu() *bodies.Menu {
	result := bodies.NewMenu()
	result.AddItemWithID("btnPreview", "#", "Preview", "fa-globe", nil)
	result.AddItemWithID("btnPublish", "#", "Publish", "fa-bolt", nil)

	return result
}

func getBlogsTopMenu() *bodies.Menu {
	result := bodies.NewMenu()
	result.AddItemWithID("btnAdd", "#", "Add Article", "fa-globe", nil)

	return result
}
