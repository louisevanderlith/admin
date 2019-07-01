package logbook

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type HistoryController struct {
	control.UIController
}

func NewHistoryCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) beego.ControllerInterface {
	result := &HistoryController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
