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

	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap, theme))

	commsCtrl := controllers.NewCommsCtrl(ctrlmap, theme)
	beego.Router("/comms/:pagesize", commsCtrl, "get:Get")
	beego.Router("/comm/:key", commsCtrl, "get:GetView")

	profileCtrl := controllers.NewProfileCtrl(ctrlmap, theme)
	beego.Router("/profiles/:pagesize", profileCtrl, "get:Get")
	beego.Router("/profile/:key", profileCtrl, "get:GetEdit")

	uploadsCtrl := controllers.NewUploadsCtrl(ctrlmap, theme)
	beego.Router("/uploads/:pagesize", uploadsCtrl, "get:Get")
	beego.Router("/upload/:key", uploadsCtrl, "get:GetView")

	beego.Router("/memory", controllers.NewMemoryCtrl(ctrlmap, theme))

	userCtrl := controllers.NewUserCtrl(ctrlmap, theme)
	beego.Router("/users/:pagesize", userCtrl, "get:Get")
	beego.Router("/user/:key", userCtrl, "get:GetView")

	commentsCtrl := controllers.NewCommentCtrl(ctrlmap, theme)
	beego.Router("/comments/:pagesize", commentsCtrl, "get:Get")
	beego.Router("/comment/:key", commentsCtrl, "get:GetView")

	vehicleCtrl := controllers.NewVehicleCtrl(ctrlmap, theme)
	beego.Router("/vehicles/:pagesize", vehicleCtrl, "get:Get")
	beego.Router("/vehicle/:key", vehicleCtrl, "get:GetView")

	vinCtrl := controllers.NewVINCtrl(ctrlmap, theme)
	beego.Router("/vins/:pagesize", vinCtrl, "get:Get")
	beego.Router("/vin/:key", vinCtrl, "get:GetView")

	blogCtrl := controllers.NewBlogCtrl(ctrlmap, theme)
	beego.Router("/blogs/:pagesize", blogCtrl, "get:Get")
	beego.Router("/blog/:key", blogCtrl, "get:GetCreate")
	beego.Router("/blog/view/:key", blogCtrl, "get:GetView")

	entityCtrl := controllers.NewEntityCtrl(ctrlmap, theme)
	beego.Router("/entities/:pagesize", entityCtrl, "get:Get")
	beego.Router("/entity/:key", entityCtrl, "get:GetEdit")

	carsCtrl := controllers.NewCarsCtrl(ctrlmap, theme)
	beego.Router("/cars/:pagesize", carsCtrl, "get:Get")
	beego.Router("/car/:key", carsCtrl, "get:GetEdit")

	partsCtrl := controllers.NewPartsCtrl(ctrlmap, theme)
	beego.Router("/parts/:pagesize", partsCtrl, "get:Get")
	beego.Router("/part/:key", partsCtrl, "get:GetEdit")

	servicesCtrl := controllers.NewServicesCtrl(ctrlmap, theme)
	beego.Router("/services/:pagesize", servicesCtrl, "get:Get")
	beego.Router("/service/:key", servicesCtrl, "get:GetEdit")
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)
	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Admin
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/comms", emptyMap)
	ctrlmap.Add("/profile", emptyMap)
	ctrlmap.Add("/memory", emptyMap)
	ctrlmap.Add("/user", emptyMap)
	ctrlmap.Add("/comment", emptyMap)
	ctrlmap.Add("/vehicle", emptyMap)
	ctrlmap.Add("/vin", emptyMap)
	ctrlmap.Add("/blog", emptyMap)
	ctrlmap.Add("/entity", emptyMap)
	ctrlmap.Add("/entities", emptyMap)
	ctrlmap.Add("/cars", emptyMap)
	ctrlmap.Add("/car", emptyMap)
	ctrlmap.Add("/services", emptyMap)
	ctrlmap.Add("/service", emptyMap)
	ctrlmap.Add("/parts", emptyMap)
	ctrlmap.Add("/part", emptyMap)


	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
