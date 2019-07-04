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
	"github.com/louisevanderlith/admin/controllers/xchange"
	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service) {
	siteProfile := beego.AppConfig.String("defaultsite")
	routes := logic.NewControlRouter(s, siteProfile)

	routes.IdentifyCtrl(controllers.NewDefaultCtrl, "", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/",
			Name:         "Home",
			Icon:         "fas fa-home",
		},
	})

	//Artifact
	routes.IdentifyCtrl(artifact.NewUploadsCtrl, "Artifact", []logic.ControlOption{
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
	})

	//Blog
	routes.IdentifyCtrl(blog.NewArticlesCtrl, "Blog", []logic.ControlOption{
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
	})

	//Comment
	routes.IdentifyCtrl(comment.NewCommentCtrl, "Comment", []logic.ControlOption{
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
	})

	//Comms
	routes.IdentifyCtrl(comms.NewMessagesCtrl, "Comms", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "message/:pagesize",
			Name:         "Messages",
			Icon:         "fas fa-envelope-square",
			Path:         "/message/:pagesize",
			Name:         "Get Messages",
			Icon:         "fa-box",
		},
		{
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/view/:key",
			Name:         "View Message",
			Icon:         "fas fa-envelope-open-text",
		},
	})

	//Entity
	routes.IdentifyCtrl(entity.NewEntitiesCtrl, "Entity", []logic.ControlOption{
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
	})

	//Folio
	routes.IdentifyCtrl(folio.NewProfileCtrl, "Folio", []logic.ControlOption{
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
	})

	//Funds
	routes.IdentifyCtrl(funds.NewAccountsCtrl, "Funds", []logic.ControlOption{
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
	})

	//Game

	//Gate

	//Logbook

	//Notify

	//Router
	routes.IdentifyCtrl(router.NewMemoryCtrl, "Router", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/memory",
			Name:         "Memory",
			Icon:         "fas fa-microchip",
		},
	})

	//Secure
	routes.IdentifyCtrl(secCtrl.NewUserCtrl, "Secure", []logic.ControlOption{
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
	})

	//Stock
	routes.IdentifyCtrl(stock.NewCarsCtrl, "Stock", []logic.ControlOption{
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
	})

	routes.IdentifyCtrl(stock.NewPartsCtrl, "Stock", []logic.ControlOption{
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
	})

	routes.IdentifyCtrl(stock.NewServicesCtrl, "Secure", []logic.ControlOption{
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
	})

	//Theme

	//Vehicle
	routes.IdentifyCtrl(vehicle.NewVehiclesCtrl, "Vehicle", []logic.ControlOption{
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
	})

	//Vin
	routes.IdentifyCtrl(vin.NewVINCtrl, "VIN", []logic.ControlOption{
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
	})

	routes.IdentifyCtrl(vin.NewRegionsCtrl, "VIN", []logic.ControlOption{
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
	})

	//XChange
	routes.IdentifyCtrl(xchange.NewCreditCtrl, "XChange", []logic.ControlOption{
		{
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/credit",
			Name:         "Credits",
			Icon:         "fas fa-hand-holding-usd",
		},
	})
}
