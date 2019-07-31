package routers

import (
	"github.com/louisevanderlith/admin/controllers"
	"github.com/louisevanderlith/admin/controllers/artifact"
	"github.com/louisevanderlith/admin/controllers/blog"
	"github.com/louisevanderlith/admin/controllers/comment"
	"github.com/louisevanderlith/admin/controllers/comms"
	"github.com/louisevanderlith/admin/controllers/curity"
	"github.com/louisevanderlith/admin/controllers/entity"
	"github.com/louisevanderlith/admin/controllers/folio"
	"github.com/louisevanderlith/admin/controllers/funds"
	"github.com/louisevanderlith/admin/controllers/router"
	"github.com/louisevanderlith/admin/controllers/stock"
	"github.com/louisevanderlith/admin/controllers/vehicle"
	"github.com/louisevanderlith/admin/controllers/vin"
	"github.com/louisevanderlith/admin/controllers/xchange"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e *droxolite.Epoxy) {
	//Home
	homeCtrl := &controllers.DefaultController{}
	homeGroup := droxolite.NewRouteGroup("", homeCtrl)
	q := make(map[string]string)
	q["access_token"] = "{access_token}"
	homeGroup.AddRouteWithQueries("/", "GET", roletype.Admin, q, homeCtrl.Get)
	homeGroup.AddRoute("/", "GET", roletype.Admin, homeCtrl.Get)
	e.AddGroup(homeGroup)

	//Artifact
	uplCtrl := &artifact.UploadsController{}
	uplGroup := droxolite.NewRouteGroup("Artifact/Uploads", uplCtrl)
	uplGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, uplCtrl.Get)
	uplGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, uplCtrl.GetView)
	e.AddGroup(uplGroup)

	//Blog
	articleCtrl := &blog.ArticlesController{}
	articleGroup := droxolite.NewRouteGroup("Blog/Articles", articleCtrl)
	articleGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, articleCtrl.Get)
	articleGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, articleCtrl.GetView)
	articleGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, articleCtrl.GetEdit)
	e.AddGroup(articleGroup)

	//Comment
	commentsCtrl := &comment.MessagesController{}
	commentsGroup := droxolite.NewRouteGroup("Comment/Messages", commentsCtrl)
	commentsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, commentsCtrl.Get)
	commentsGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, commentsCtrl.GetView)
	e.AddGroup(commentsGroup)

	//Comms
	messageCtrl := &comms.MessagesController{}
	messageGroup := droxolite.NewRouteGroup("Comms/Messages", messageCtrl)
	messageGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, messageCtrl.Get)
	messageGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, messageCtrl.GetView)
	e.AddGroup(messageGroup)

	//Entity
	entityCtrl := &entity.EntitiesController{}
	entityGroup := droxolite.NewRouteGroup("Entity/Entities", entityCtrl)
	entityGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, entityCtrl.Get)
	entityGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, entityCtrl.GetEdit)
	e.AddGroup(entityGroup)

	//Folio
	profileCtrl := &folio.ProfileController{}
	profileGroup := droxolite.NewRouteGroup("Folio/Profiles", profileCtrl)
	profileGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, profileCtrl.Get)
	profileGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, profileCtrl.GetEdit)
	e.AddGroup(profileGroup)

	//Funds
	accountCtrl := &funds.AccountsController{}
	accountGroup := droxolite.NewRouteGroup("Funds/Accounts", accountCtrl)
	accountGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, accountCtrl.Get)
	accountGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, accountCtrl.GetEdit)
	e.AddGroup(accountGroup)

	//Game

	//Gate

	//Logbook

	//Notify

	//Router
	memCtrl := &router.MemoryController{}
	memGroup := droxolite.NewRouteGroup("Router/Memory", memCtrl)
	memGroup.AddRoute("/", "GET", roletype.Admin, memCtrl.Get)
	e.AddGroup(memGroup)

	//Secure
	userCtrl := &curity.UserController{}
	userGroup := droxolite.NewRouteGroup("Secure/Users", userCtrl)
	userGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, userCtrl.Get)
	userGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, userCtrl.GetView)
	e.AddGroup(userGroup)

	//Stock
	carsCtrl := &stock.CarsController{}
	carsGroup := droxolite.NewRouteGroup("Stock/Cars", carsCtrl)
	carsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, carsCtrl.Get)
	carsGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, carsCtrl.GetEdit)
	e.AddGroup(carsGroup)

	partsCtrl := &stock.PartsController{}
	partsGroup := droxolite.NewRouteGroup("Stock/Parts", partsCtrl)
	partsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, partsCtrl.Get)
	partsGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, partsCtrl.GetEdit)
	e.AddGroup(partsGroup)

	srvcCtrl := &stock.ServicesController{}
	srvcGroup := droxolite.NewRouteGroup("Stock/Services", srvcCtrl)
	srvcGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, srvcCtrl.Get)
	srvcGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, srvcCtrl.GetEdit)
	e.AddGroup(srvcGroup)

	//Theme

	//Vehicle
	vehsCtrl := &vehicle.VehiclesController{}
	vehsGroup := droxolite.NewRouteGroup("Vehicle/Vehicles", vehsCtrl)
	vehsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, vehsCtrl.Get)
	vehsGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, vehsCtrl.GetView)
	e.AddGroup(vehsGroup)

	//Vin
	vinCtrl := &vin.VINController{}
	vinGroup := droxolite.NewRouteGroup("VIN/VINS", vinCtrl)
	vinGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, vinCtrl.Get)
	vinGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, vinCtrl.GetView)
	e.AddGroup(vehsGroup)

	regionCtrl := &vin.RegionsController{}
	regionGroup := droxolite.NewRouteGroup("VIN/Regions", regionCtrl)
	regionGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, regionCtrl.Get)
	regionGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, regionCtrl.GetEdit)
	e.AddGroup(regionGroup)

	//XChange
	creditCtrl := &xchange.CreditsController{}
	creditGroup := droxolite.NewRouteGroup("XChange/Credits", creditCtrl)
	creditGroup.AddRoute("/", "GET", roletype.Admin, creditCtrl.Get)
	e.AddGroup(creditGroup)
}
