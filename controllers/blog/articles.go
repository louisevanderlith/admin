package blog

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Articles struct {
}

func (c *Articles) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("articles", "Articles", true)

	//ctx.CreateTopMenu(getBlogsTopMenu())

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "article", "all", "non", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Articles) Search(ctx context.Contexer) (int, interface{}) {
	//c.Setup("articles", "Articles", true)

	//c.CreateTopMenu(getBlogsTopMenu())

	var result []interface{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "article", "all", "non", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Articles) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("articleEdit", "Edit Article", true)

	//c.CreateTopMenu(createBlogTopMenu())
	//c.EnableSave()

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Blog.API", "article", key.String())

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
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
