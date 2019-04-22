package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/admin/controllers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilters(s)

	siteName := beego.AppConfig.String("defaultsite")
	theme, err := mango.GetDefaultTheme(ctrlmap.GetInstanceID(), siteName)

	if err != nil {
		panic(err)
	}

	commsCtrl := controllers.NewCommsCtrl(ctrlmap, theme)
	profileCtrl := controllers.NewProfileCtrl(ctrlmap, theme)

	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap, theme))
	beego.Router("/comms", commsCtrl, "get:Get")
	beego.Router("/comms/:key", commsCtrl, "get:GetView")
	beego.Router("/profiles/:pagesize", profileCtrl, "get:Get")
	beego.Router("/profile/:key", profileCtrl, "get:GetEdit")
	beego.Router("/memory", controllers.NewMemoryCtrl(ctrlmap, theme))
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)
	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Admin
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/comms", emptyMap)
	ctrlmap.Add("/profiles", emptyMap)
	ctrlmap.Add("/profile", emptyMap)
	ctrlmap.Add("/memory", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
