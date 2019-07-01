package logic

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

type ControlRouter struct {
	Menu        *control.Menu
	Settings    mango.ThemeSetting
	Mapping     *control.ControllerMap
	Controllers map[string]*control.UIController
}

func NewControlRouter(servc *mango.Service, siteProfile string) *ControlRouter {
	mapping := control.CreateControlMap(servc)
	setting, err := mango.GetDefaultTheme(mapping.GetInstanceID(), siteProfile)

	if err != nil {
		panic(err)
	}

	routes := &ControlRouter{
		Menu:     control.NewMenu("/"),
		Settings: setting,
		Mapping:  mapping,
	}

	beego.InsertFilter("/*", beego.BeforeRouter, routes.SetActivePath)
	beego.InsertFilter("/*", beego.BeforeRouter, mapping.FilterUI)

	return routes
}

type ControlOption struct {
	Method       string
	RequiredRole roletype.Enum
	Function     string
	Path         string
	Name         string
	Icon         string
}

type ControlConstructor func(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) beego.ControllerInterface

func (r *ControlRouter) SetActivePath(ctx *context.Context) {
	path := ctx.Request.URL.RequestURI()
	r.Menu.SetActive(path)

	for p, ctrl := range r.Controllers {
		if strings.HasPrefix(path, p) {
			ctrl.CreateSideMenu(r.Menu)
		}
	}
}

func (r *ControlRouter) IdentifyCtrl(ctor ControlConstructor, name string, options []ControlOption) {
	ctrllr := ctor(r.Mapping, r.Settings)

	actMap := make(secure.ActionMap)
	basePath := fmt.Sprintf("/%s/", name)

	r.Controllers[basePath] = ctrllr.(*control.UIController)

	children := control.NewMenu(basePath)

	for _, v := range options {
		actMap[v.Method] = v.RequiredRole
		funcPath := basePath + v.Path

		r.Mapping.Add(funcPath, actMap)

		realPath := strings.Replace(funcPath, ":pagesize", "A10", 1)

		if !strings.Contains(v.Path, ":key") {
			children.AddItem(realPath, v.Name, v.Icon, nil)
		}

		actFunc := fmt.Sprintf("%s:%s", strings.ToLower(v.Method), v.Function)
		beego.Router(funcPath, ctrllr, actFunc)
	}

	r.Menu.AddItem(basePath, name, "", children)
}
