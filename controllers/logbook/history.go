package logbook

import (
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type HistoryController struct {
	control.UIController
}

func NewHistoryCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) logic.PageUI {
	result := &HistoryController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
