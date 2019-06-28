package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/admin/controllers"
	"github.com/louisevanderlith/admin/controllers/artifact"
	"github.com/louisevanderlith/admin/controllers/blog"
	"github.com/louisevanderlith/admin/controllers/comment"
	"github.com/louisevanderlith/admin/controllers/comms"
	"github.com/louisevanderlith/admin/controllers/entity"
	"github.com/louisevanderlith/admin/controllers/folio"
	"github.com/louisevanderlith/admin/controllers/funds"
	"github.com/louisevanderlith/admin/controllers/router"
	secCtrl "github.com/louisevanderlith/admin/controllers/secure"
	"github.com/louisevanderlith/admin/controllers/stock"
	"github.com/louisevanderlith/admin/controllers/vehicle"
	"github.com/louisevanderlith/admin/controllers/vin"
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

	//Artifact
	uploadsCtrl := artifact.NewUploadsCtrl(ctrlmap, theme)
	beego.Router("/artifact/uploads/:pagesize", uploadsCtrl, "get:Get")
	beego.Router("/artifact/upload/:key", uploadsCtrl, "get:GetView")

	//Blog
	articlesCtrl := blog.NewArticlesCtrl(ctrlmap, theme)
	beego.Router("/blog/articles/:pagesize", articlesCtrl, "get:Get")
	beego.Router("/blog/article/:key", articlesCtrl, "get:GetCreate")
	beego.Router("/blog/view/:key", articlesCtrl, "get:GetView")

	//Comment
	commentsCtrl := comment.NewCommentCtrl(ctrlmap, theme)
	beego.Router("/comment/comments/:pagesize", commentsCtrl, "get:Get")
	beego.Router("/comment/comment/:key", commentsCtrl, "get:GetView")

	//Comms
	commsCtrl := comms.NewMessagesCtrl(ctrlmap, theme)
	beego.Router("/comms/messages/:pagesize", commsCtrl, "get:Get")
	beego.Router("/comms/message/:key", commsCtrl, "get:GetView")

	//Entity
	entityCtrl := entity.NewEntitiesCtrl(ctrlmap, theme)
	beego.Router("/entity/entities/:pagesize", entityCtrl, "get:Get")
	beego.Router("/entity/view/:key", entityCtrl, "get:GetEdit")

	//Folio
	profileCtrl := folio.NewProfileCtrl(ctrlmap, theme)
	beego.Router("/folio/profiles/:pagesize", profileCtrl, "get:Get")
	beego.Router("/folio/profile/:key", profileCtrl, "get:GetEdit")

	//Funds
	accountCtrl := funds.NewAccountsCtrl(ctrlmap, theme)
	beego.Router("/funds/accounts/:pagesize", accountCtrl, "get:Get")
	beego.Router("/funds/account/:key", accountCtrl, "get:GetEdit")

	//Game

	//Gate

	//Logbook

	//Notify

	//Router
	beego.Router("/router/memory", router.NewMemoryCtrl(ctrlmap, theme))

	//Secure
	userCtrl := secCtrl.NewUserCtrl(ctrlmap, theme)
	beego.Router("/secure/users/:pagesize", userCtrl, "get:Get")
	beego.Router("/secure/user/:key", userCtrl, "get:GetView")

	//Stock
	carsCtrl := stock.NewCarsCtrl(ctrlmap, theme)
	beego.Router("/stock/cars/:pagesize", carsCtrl, "get:Get")
	beego.Router("/stock/car/:key", carsCtrl, "get:GetEdit")

	partsCtrl := stock.NewPartsCtrl(ctrlmap, theme)
	beego.Router("/stock/parts/:pagesize", partsCtrl, "get:Get")
	beego.Router("/stock/part/:key", partsCtrl, "get:GetEdit")

	servicesCtrl := stock.NewServicesCtrl(ctrlmap, theme)
	beego.Router("/stock/services/:pagesize", servicesCtrl, "get:Get")
	beego.Router("/stock/service/:key", servicesCtrl, "get:GetEdit")

	//Theme

	//Vehicle
	vehicleCtrl := vehicle.NewVehiclesCtrl(ctrlmap, theme)
	beego.Router("/vehicles/:pagesize", vehicleCtrl, "get:Get")
	beego.Router("/vehicle/:key", vehicleCtrl, "get:GetView")

	//Vin
	vinCtrl := vin.NewVINCtrl(ctrlmap, theme)
	beego.Router("/vin/vins/:pagesize", vinCtrl, "get:Get")
	beego.Router("/vin/vin/:key", vinCtrl, "get:GetView")

	regnCtrl := vin.NewRegionsCtrl(ctrlmap, theme)
	beego.Router("/vin/regions/:pagesize", regnCtrl, "get:Get")
	beego.Router("/vin/region/:key", regnCtrl, "get:GetEdit")

	//XChange

}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)
	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Admin
	emptyMap["GET"] = roletype.Admin

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/artifact/upload", emptyMap)
	ctrlmap.Add("/blog/article", emptyMap)
	ctrlmap.Add("/blog/view", emptyMap)
	ctrlmap.Add("/comment/comment", emptyMap)
	ctrlmap.Add("/comms/message", emptyMap)
	ctrlmap.Add("/entity/entities", emptyMap)
	ctrlmap.Add("/entity/entity", emptyMap)
	ctrlmap.Add("/profile", emptyMap)
	ctrlmap.Add("/memory", emptyMap)
	ctrlmap.Add("/user", emptyMap)
	ctrlmap.Add("/comment", emptyMap)
	ctrlmap.Add("/vehicle", emptyMap)
	ctrlmap.Add("/vin", emptyMap)

	ctrlmap.Add("/accounts", emptyMap)
	ctrlmap.Add("/account", emptyMap)

	ctrlmap.Add("/history", emptyMap)

	ctrlmap.Add("/cars", emptyMap)
	ctrlmap.Add("/car", emptyMap)
	ctrlmap.Add("/services", emptyMap)
	ctrlmap.Add("/service", emptyMap)
	ctrlmap.Add("/parts", emptyMap)
	ctrlmap.Add("/part", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
