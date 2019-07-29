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
	"github.com/louisevanderlith/droxolite/roletype"
)

type PageUI interface {
	xontrols.Controller
	CreateTopMenu(menu *bodies.Menu)
	CreateSideMenu(menu *bodies.Menu)
}

type CtrlSet struct {
	UI   PageUI
	Path string
}

type ControlRouter struct {
	Menu        *bodies.Menu
	Settings    droxolite.ThemeSetting
	Mapping     *control.ControllerMap
	Controllers []CtrlSet
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

	beego.InsertFilter("/*", beego.BeforeRouter, mapping.FilterUI)
	beego.InsertFilter("/*", beego.BeforeExec, routes.SetActivePath)

	return routes
}

type ControlOption struct {
	RequiredRole roletype.Enum
	Function     string
	Path         string
	Name         string
	Icon         string
}

type ControlConstructor func(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) PageUI

func (r *ControlRouter) SetActivePath(ctx *context.Context) {
	path := strings.ToLower(ctx.Request.URL.RequestURI())
	r.Menu.SetActive(path)
	
	for _, ctrl := range r.Controllers {

		log.Printf("Matching %s vs %s\n", path, ctrl.Path)

		if ctrl.Path == "" && path == "/" {
			ctrl.UI.CreateSideMenu(ctx, r.Menu)
			return
		}

		if len(ctrl.Path) > 1 && strings.HasPrefix(path, ctrl.Path) {
			log.Print("passed")
			ctrl.UI.CreateSideMenu(ctx, r.Menu)
		}
	}
}

func (r *ControlRouter) IdentifyCtrl(ctor ControlConstructor, name string, options []ControlOption) {
	instance := ctor(r.Mapping, r.Settings)

	actMap := make(secure.ActionMap)
	basePath := ""

	if len(name) != 0 {
		basePath = strings.ToLower(fmt.Sprintf("/%s", name))
	}

	r.Controllers = append(r.Controllers, CtrlSet{UI: instance, Path: basePath})

	children := control.NewMenu(basePath)

	for _, v := range options {
		actMap["GET"] = v.RequiredRole
		funcPath := strings.ToLower(basePath + v.Path)

		r.Mapping.Add(funcPath, actMap)

		realPath := strings.Replace(funcPath, ":pagesize", "A10", 1)

		if !strings.Contains(strings.ToLower(v.Path), ":key") {
			children.AddItem(realPath, v.Name, v.Icon, nil)
		}

		actFunc := fmt.Sprintf("get:%s", v.Function)
		log.Println(funcPath)
		beego.Router(funcPath, instance, actFunc)
	}

	r.Menu.AddItem(basePath, name, "", children)
}
