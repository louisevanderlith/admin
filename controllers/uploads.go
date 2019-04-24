package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type UploadsController struct {
	control.UIController
}

func NewUploadsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *UploadsController {
	result := &UploadsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
