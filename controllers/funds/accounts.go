package funds

import (
	"log"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"

	"github.com/louisevanderlith/mango/control"

	"github.com/louisevanderlith/admin/logic"
)

type AccountsController struct {
	control.UIController
}

func NewAccountsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *AccountsController {
	result := &AccountsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *AccountsController) Get() {
	c.Setup("accounts", "Accounts", true)
	c.CreateSideMenu(logic.GetMenu("/accounts"))

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Funds.API", "account", "all", pagesize)

	c.Serve(result, err)
}

func (c *AccountsController) GetEdit() {
	c.Setup("accountView", "View Account", true)
	c.CreateSideMenu(logic.GetMenu("/accounts"))
	c.EnableSave()

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Funds.API", "account", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
