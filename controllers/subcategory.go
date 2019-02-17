package controllers

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango/control"
)

type SubCategoryController struct {
	control.UIController
}

func NewSubCategoryCtrl(ctrlMap *control.ControllerMap) *SubCategoryController {
	result := &SubCategoryController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *SubCategoryController) Get() {
	c.Setup("subcategory")
	c.CreateSideMenu(logic.GetMenu("/subcategory"))

	data, err := logic.GetSubCategories(c.GetInstanceID())

	c.Serve(data, err)
}