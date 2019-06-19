package xchange

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type CreditsController struct {
	control.UIController
}

func NewCreditCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *CreditsController {
	result := &CreditsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *CreditsController) Get() {
	c.Setup("credits", "Credits", true)
	c.CreateSideMenu(logic.GetMenu("/credits"))

	c.Serve(nil, nil)
}

func (c *CreditsController) GetView() {
	c.Setup("creditView", "View Credit", false)
	c.CreateSideMenu(logic.GetMenu("/credit"))

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "XChange.API", "???", key.String())

	c.Serve(result, err)
}
