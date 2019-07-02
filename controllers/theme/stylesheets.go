package theme

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type StylesheetController struct {
	control.UIController
}

func NewStylesheetCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) logic.PageUI {
	result := &StylesheetController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *StylesheetController) Get() {
	c.Setup("stylesheets", "Stylesheets", false)

	result := []interface{}{}
	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Theme.API", "asset", "css")

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(nil, nil)
}

func (c *StylesheetController) GetView() {
	c.Setup("stylesheetView", "View Stylesheet", false)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Theme.API", "???", key.String())

	c.Serve(result, err)
}
