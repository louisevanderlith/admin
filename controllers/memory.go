package controllers

import (
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

	result := []interface{}{}
	err := mango.DoGET(&result, c.GetInstanceID(), "Router.API", "memory")

	c.Serve(result, err)
}
