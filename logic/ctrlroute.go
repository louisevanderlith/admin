package logic

import (
	"fmt"
	"log"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

type PageUI interface {
	beego.ControllerInterface
	CreateTopMenu(ctx *context.Context, menu *control.Menu)
	CreateSideMenu(ctx *context.Context, menu *control.Menu)
}

type ControlRouter struct {
	Menu        *control.Menu
	Settings    mango.ThemeSetting
	Mapping     *control.ControllerMap
	Controllers map[string]PageUI
}

func NewControlRouter(servc *mango.Service, siteProfile string) *ControlRouter {
	mapping := control.CreateControlMap(servc)
	setting, err := mango.GetDefaultTheme(mapping.GetInstanceID(), siteProfile)

	if err != nil {
		panic(err)
	}

	routes := &ControlRouter{
		Menu:        control.NewMenu("/"),
		Settings:    setting,
		Mapping:     mapping,
		Controllers: make(map[string]PageUI),
	}

	beego.InsertFilter("/*", beego.BeforeRouter, mapping.FilterUI)
	beego.InsertFilter("/*", beego.BeforeExec, routes.SetActivePath)

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

type ControlConstructor func(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) PageUI

func (r *ControlRouter) SetActivePath(ctx *context.Context) {
	path := ctx.Request.URL.RequestURI()
	r.Menu.SetActive(path)
	log.Print("Setting Active Path")
	for p, ctrl := range r.Controllers {
		log.Printf("Matching %s vs %s\n", path, p)
		if strings.HasPrefix(path, p) {
			log.Print("passed")
			ctrl.CreateSideMenu(ctx, r.Menu)
		}
	}
}

func (r *ControlRouter) IdentifyCtrl(ctor ControlConstructor, name string, options []ControlOption) {
	ctrllr := ctor(r.Mapping, r.Settings)

	actMap := make(secure.ActionMap)
	basePath := ""

	if len(name) != 0 {
		basePath = fmt.Sprintf("/%s", name)
	}

	r.Controllers[basePath] = ctrllr

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
