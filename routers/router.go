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
	homeGroup.AddRoute("/", "GET", roletype.Admin, homeCtrl.Get)
	e.AddGroup(homeGroup)

	/*
		routes := logic.NewControlRouter(s, profile)

		routes.IdentifyCtrl(controllers.NewDefaultCtrl, "", []logic.ControlOption{
			{
				RequiredRole: roletype.Admin,
				Function:     "Get",
				Path:         "/",
				Name:         "Home",
				Icon:         "fas fa-home",
			},
		})
	*/
	//Artifact
	uplCtrl := &artifact.UploadsController{}
	uplGroup := droxolite.NewRouteGroup("Artifact/Uploads", uplCtrl)
	uplGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, uplCtrl.Get)
	uplGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, uplCtrl.GetView)
	e.AddGroup(uplGroup)
	/*routes.IdentifyCtrl(artifact.NewUploadsCtrl, "Artifact", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/uploads/:pagesize",
			Name:         "Uploads",
			Icon:         "fas fa-box-open",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/upload/:key",
			Name:         "View Upload",
			Icon:         "fas fa-images",
		},
	})*/

	//Blog
	articleCtrl := &blog.ArticlesController{}
	articleGroup := droxolite.NewRouteGroup("Blog/Articles", articleCtrl)
	articleGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, articleCtrl.Get)
	articleGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, articleCtrl.GetView)
	articleGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, articleCtrl.GetEdit)
	e.AddGroup(articleGroup)
	/*routes.IdentifyCtrl(blog.NewArticlesCtrl, "Blog", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/articles/:pagesize",
			Name:         "Articles",
			Icon:         "fas fa-box-open",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/view/:key",
			Name:         "View Article",
			Icon:         "fas fa-images",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetCreate",
			Path:         "/article/:key",
			Name:         "Edit Article",
			Icon:         "fas fa-edit",
		},
	})*/

	//Comment
	commentsCtrl := &comment.MessagesController{}
	commentsGroup := droxolite.NewRouteGroup("Comment/Messages", commentsCtrl)
	commentsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, commentsCtrl.Get)
	commentsGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, commentsCtrl.GetView)
	e.AddGroup(commentsGroup)
	/*.IdentifyCtrl(comment.NewCommentCtrl, "Comment", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/comments/:pagesize",
			Name:         "Comments",
			Icon:         "fas fa-inbox",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/view/:key",
			Name:         "View Comment",
			Icon:         "fas fa-comments",
		},
	})*/

	//Comms
	messageCtrl := &comms.MessagesController{}
	messageGroup := droxolite.NewRouteGroup("Comms/Messages", messageCtrl)
	messageGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, messageCtrl.Get)
	messageGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, messageCtrl.GetView)
	e.AddGroup(messageGroup)

	/*routes.IdentifyCtrl(comms.NewMessagesCtrl, "Comms", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/message/:pagesize",
			Name:         "Messages",
			Icon:         "fas fa-envelope-square",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/view/:key",
			Name:         "View Message",
			Icon:         "fas fa-envelope-open-text",
		},
	})*/

	//Entity
	entityCtrl := &entity.EntitiesController{}
	entityGroup := droxolite.NewRouteGroup("Entity/Entities", entityCtrl)
	entityGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, entityCtrl.Get)
	entityGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, entityCtrl.GetEdit)
	e.AddGroup(entityGroup)

	/*routes.IdentifyCtrl(entity.NewEntitiesCtrl, "Entity", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/entities/:pagesize",
			Name:         "Entities",
			Icon:         "fas fa-users",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/entity/:key",
			Name:         "Edit Entity",
			Icon:         "fas fa-house-damage",
		},
	})*/

	//Folio
	profileCtrl := &folio.ProfileController{}
	profileGroup := droxolite.NewRouteGroup("Folio/Profiles", profileCtrl)
	profileGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, profileCtrl.Get)
	profileGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, profileCtrl.GetEdit)
	e.AddGroup(profileGroup)
	/*routes.IdentifyCtrl(folio.NewProfileCtrl, "Folio", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/profiles/:pagesize",
			Name:         "Profiles",
			Icon:         "fas fa-id-card-alt",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/profile/:key",
			Name:         "Edit Profile",
			Icon:         "fas fa-user-circle",
		},
	})*/

	//Funds
	accountCtrl := &funds.AccountsController{}
	accountGroup := droxolite.NewRouteGroup("Funds/Accounts", accountCtrl)
	accountGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, accountCtrl.Get)
	accountGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, accountCtrl.GetEdit)
	e.AddGroup(accountGroup)
	/*routes.IdentifyCtrl(funds.NewAccountsCtrl, "Funds", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/accounts/:pagesize",
			Name:         "Accounts",
			Icon:         "fas fa-file-invoice",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/account/:key",
			Name:         "Edit Account",
			Icon:         "fas fa-file-invoice",
		},
	})*/

	//Game

	//Gate

	//Logbook

	//Notify

	//Router
	memCtrl := &router.MemoryController{}
	memGroup := droxolite.NewRouteGroup("Router/Memory", memCtrl)
	memGroup.AddRoute("/", "GET", roletype.Admin, memCtrl.Get)
	e.AddGroup(memGroup)

	/*routes.IdentifyCtrl(router.NewMemoryCtrl, "Router", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/memory",
			Name:         "Memory",
			Icon:         "fas fa-microchip",
		},
	})*/

	//Secure
	userCtrl := &curity.UserController{}
	userGroup := droxolite.NewRouteGroup("Secure/Users", userCtrl)
	userGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, userCtrl.Get)
	userGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, userCtrl.GetView)
	e.AddGroup(userGroup)
	/*routes.IdentifyCtrl(secCtrl.NewUserCtrl, "Secure", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/users/:pagesize",
			Name:         "Users",
			Icon:         "fas fa-user-friends",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/user/:key",
			Name:         "View User",
			Icon:         "fas fa-street-view",
		},
	})*/

	//Stock
	carsCtrl := &stock.CarsController{}
	carsGroup := droxolite.NewRouteGroup("Stock/Cars", carsCtrl)
	carsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, carsCtrl.Get)
	carsGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, carsCtrl.GetEdit)
	e.AddGroup(carsGroup)
	/*routes.IdentifyCtrl(stock.NewCarsCtrl, "Stock", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/cars/:pagesize",
			Name:         "Cars",
			Icon:         "fas fa-car-side",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/car/:key",
			Name:         "Car",
			Icon:         "fas fa-car-crash",
		},
	})*/

	partsCtrl := &stock.PartsController{}
	partsGroup := droxolite.NewRouteGroup("Stock/Parts", partsCtrl)
	partsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, partsCtrl.Get)
	partsGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, partsCtrl.GetEdit)
	e.AddGroup(partsGroup)

	/*routes.IdentifyCtrl(stock.NewPartsCtrl, "Stock", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/parts/:pagesize",
			Name:         "Parts",
			Icon:         "fas fa-wrench",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/part/:key",
			Name:         "Edit Part",
			Icon:         "fas fa-cog",
		},
	})*/

	srvcCtrl := &stock.ServicesController{}
	srvcGroup := droxolite.NewRouteGroup("Stock/Services", srvcCtrl)
	srvcGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, srvcCtrl.Get)
	srvcGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, srvcCtrl.GetEdit)
	e.AddGroup(srvcGroup)
	/*routes.IdentifyCtrl(stock.NewServicesCtrl, "Stock", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/services/:pagesize",
			Name:         "Services",
			Icon:         "fas fa-layer-group",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/service/:key",
			Name:         "Edit Service",
			Icon:         "fas fa-edit",
		},
	})*/

	//Theme

	//Vehicle
	vehsCtrl := &vehicle.VehiclesController{}
	vehsGroup := droxolite.NewRouteGroup("Vehicle/Vehicles", vehsCtrl)
	vehsGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, vehsCtrl.Get)
	vehsGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, vehsCtrl.GetView)
	e.AddGroup(vehsGroup)
	/*routes.IdentifyCtrl(vehicle.NewVehiclesCtrl, "Vehicle", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/vehicles/:pagesize",
			Name:         "Vehicles",
			Icon:         "fas fa-truck-pickup",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/vehicle/:key",
			Name:         "View Vehicle",
			Icon:         "fas fa-car",
		},
	})*/

	//Vin
	vinCtrl := &vin.VINController{}
	vinGroup := droxolite.NewRouteGroup("VIN/VINS", vinCtrl)
	vinGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, vinCtrl.Get)
	vinGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, vinCtrl.GetView)
	e.AddGroup(vehsGroup)
	/*routes.IdentifyCtrl(vin.NewVINCtrl, "VIN", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/vins/:pagesize",
			Name:         "VINs",
			Icon:         "fas fa-barcode",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/vin/:key",
			Name:         "View VIN",
			Icon:         "fa-binoculers",
		},
	})*/

	regionCtrl := &vin.RegionsController{}
	regionGroup := droxolite.NewRouteGroup("VIN/Regions", regionCtrl)
	regionGroup.AddRoute("/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, regionCtrl.Get)
	regionGroup.AddRoute("/edit/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, regionCtrl.GetEdit)
	e.AddGroup(regionGroup)
	/*routes.IdentifyCtrl(vin.NewRegionsCtrl, "VIN", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/regions/:pagesize",
			Name:         "Regions",
			Icon:         "fas fa-globe-africa",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/region/:key",
			Name:         "Region",
			Icon:         "fas fa-flag",
		},
	})*/

	//XChange
	creditCtrl := &xchange.CreditsController{}
	creditGroup := droxolite.NewRouteGroup("XChange/Credits", creditCtrl)
	creditGroup.AddRoute("/", "GET", roletype.Admin, creditCtrl.Get)
	e.AddGroup(creditGroup)
	/*routes.IdentifyCtrl(xchange.NewCreditCtrl, "XChange", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/credit",
			Name:         "Credits",
			Icon:         "fas fa-hand-holding-usd",
		},
	})*/
}
