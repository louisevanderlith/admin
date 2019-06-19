package logbook

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type NotificationsController struct {
	control.UIController
}

func NewNotficationsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *NotificationsController {
	result := &NotificationsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
