package theme

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type TemplatesController struct {
	control.UIController
}

func NewTemplatesCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) logic.PageUI {
	result := &TemplatesController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *TemplatesController) Get() {
	c.Setup("templates", "Templates", false)

	result := []interface{}{}
	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Theme.API", "asset", "html")

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(nil, nil)
}

func (c *TemplatesController) GetView() {
	c.Setup("templateView", "View Template", false)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Theme.API", "???", key.String())

	c.Serve(result, err)
}
