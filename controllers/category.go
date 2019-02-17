package controllers

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango/control"
)

type CategoryController struct {
	control.UIController
}

func NewCategoryCtrl(ctrlMap *control.ControllerMap) *CategoryController {
	result := &CategoryController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *CategoryController) Get() {
	c.Setup("category")
	c.CreateSideMenu(logic.GetMenu("/category"))

	data, err := logic.GetCategories(c.GetInstanceID())

	c.Serve(data, err)
}
