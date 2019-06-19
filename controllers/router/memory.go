package router

import (
	"log"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type MemoryController struct {
	control.UIController
}

func NewMemoryCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *MemoryController {
	result := &MemoryController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *MemoryController) Get() {
	c.Setup("memory", "Memory", true)
	c.CreateSideMenu(logic.GetMenu("/memory"))

	result := make(map[string]interface{})
	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Router.API", "memory")

	if err != nil {
		log.Println(err)
		c.Serve(nil, err)
		return
	}

	c.Serve(result, nil)
}
