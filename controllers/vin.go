package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type VINController struct {
	control.UIController
}

func NewVINCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *VINController {
	result := &VINController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
