package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type CommentController struct {
	control.UIController
}

func NewCommentCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *CommentController {
	result := &CommentController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}
