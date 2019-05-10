package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type VehicleController struct {
	control.UIController
}

func NewVehicleCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *VehicleController {
	result := &VehicleController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
