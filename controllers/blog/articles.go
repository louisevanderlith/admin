package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Articles struct {
	xontrols.UICtrl
}

func (c *Articles) Default() {
	c.Setup("articles", "Articles", true)

	c.CreateTopMenu(getBlogsTopMenu())

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", "all", "non", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Articles) Search() {
	c.Setup("articles", "Articles", true)

	c.CreateTopMenu(getBlogsTopMenu())

	var result []interface{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", "all", "non", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Articles) View() {
	c.Setup("articleView", "View Article", true)

	c.CreateTopMenu(createBlogTopMenu())
	c.EnableSave()

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	var result []interface{}
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

/*
func (c *Articles) Create() {
	c.Setup("articleCreate", "Create Article", true)

	c.CreateTopMenu(createBlogTopMenu())
	c.EnableSave()

	err := c.Serve(http.StatusOK, nil, nil)

	if err != nil {
		log.Println(err)
	}
}*/

func createBlogTopMenu() bodies.Menu {
	result := bodies.NewMenu()

	items := []bodies.MenuItem{
		bodies.NewItem("btnPreview", "#", "Preview", nil),
		bodies.NewItem("btnPublish", "#", "Publish", nil),
	}

	result.AddGroup("Buttons", items)

	return *result
}

func getBlogsTopMenu() bodies.Menu {
	result := bodies.NewMenu()
	item := bodies.NewItem("btnAdd", "", "Add Article", nil)
	result.AddGroup("Buttons", []bodies.MenuItem{item})

	return *result
}
