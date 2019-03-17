package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/admin/controllers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/mango/enums"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilters(s)

	commsCtrl := controllers.NewCommsCtrl(ctrlmap)
	profileCtrl := controllers.NewProfileCtrl(ctrlmap)

	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap))
	beego.Router("/category", controllers.NewCategoryCtrl(ctrlmap))
	beego.Router("/comms", commsCtrl, "get:Get")
	beego.Router("/comms/:key", commsCtrl, "get:GetView")
	beego.Router("/manufacturer", controllers.NewManufacturerCtrl(ctrlmap))
	beego.Router("/model", controllers.NewModelCtrl(ctrlmap))
	beego.Router("/subcategory", controllers.NewSubCategoryCtrl(ctrlmap))
	beego.Router("/profiles", profileCtrl, "get:Get")
	beego.Router("/profile/:key", profileCtrl, "get:GetEdit")
	beego.Router("/memory", controllers.NewMemoryCtrl(ctrlmap))
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)
	emptyMap := make(control.ActionMap)
	emptyMap["POST"] = enums.Admin
	emptyMap["GET"] = enums.Admin

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/category", emptyMap)
	ctrlmap.Add("/comms", emptyMap)
	ctrlmap.Add("/manufacturer", emptyMap)
	ctrlmap.Add("/model", emptyMap)
	ctrlmap.Add("/subcategory", emptyMap)
	ctrlmap.Add("/profiles", emptyMap)
	ctrlmap.Add("/profile", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
