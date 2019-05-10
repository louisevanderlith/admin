package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type HeroController struct {
	control.UIController
}

func NewHeroCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *HeroController {
	result := &HeroController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
