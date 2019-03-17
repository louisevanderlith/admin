package controllers

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type MemoryController struct {
	control.UIController
}

func NewMemoryCtrl(ctrlMap *control.ControllerMap) *MemoryController {
	result := &MemoryController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *MemoryController) Get() {
	c.Setup("memory")
	c.CreateSideMenu(logic.GetMenu("/memory"))

	result := []interface{}{}
	fail, err := mango.DoGET(&result, c.GetInstanceID(), "Router.API", "memory")

	if err != nil {
		c.Serve(nil, err)
		return
	}

	if fail != nil {
		c.Serve(nil, fail)
		return
	}

	c.Serve(result, nil)
}
