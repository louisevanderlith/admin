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

	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap))
	beego.Router("/category", controllers.NewCategoryCtrl(ctrlmap))
	beego.Router("/comms", controllers.NewCommsCtrl(ctrlmap))
	beego.Router("/manufacturer", controllers.NewManufacturerCtrl(ctrlmap))
	beego.Router("/model", controllers.NewModelCtrl(ctrlmap))
	beego.Router("/subcategory", controllers.NewSubCategoryCtrl(ctrlmap))
	beego.Router("/profiles", controllers.NewProfileCtrl(ctrlmap))
	beego.Router("/profile/:key", controllers.NewProfileCtrl(ctrlmap), "get:GetEdit")
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

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
