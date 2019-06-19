package logbook

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type HistoryController struct {
	control.UIController
}

func NewHistoryCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *HistoryController {
	result := &HistoryController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
