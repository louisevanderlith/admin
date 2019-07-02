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
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/",
			Name:         "Home",
			Icon:         "fa-home",
		},
	})

	//Artifact
	routes.IdentifyCtrl(artifact.NewUploadsCtrl, "Artifact", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/uploads/:pagesize",
			Name:         "Get Uploads",
			Icon:         "fa-box",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/upload/:key",
			Name:         "View Upload",
			Icon:         "fa-image",
		},
	})

	//Blog
	routes.IdentifyCtrl(blog.NewArticlesCtrl, "Blog", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/articles/:pagesize",
			Name:         "Get Articles",
			Icon:         "fa-box",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/view/:key",
			Name:         "View Article",
			Icon:         "fa-image",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetCreate",
			Path:         "/article/:key",
			Name:         "Edit Article",
			Icon:         "fa-image",
		},
	})

	//Comment
	routes.IdentifyCtrl(comment.NewCommentCtrl, "Comment", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/comments/:pagesize",
			Name:         "Get Comments",
			Icon:         "fa-box",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/view/:key",
			Name:         "View Comment",
			Icon:         "fa-image",
		},
	})

	//Comms
	routes.IdentifyCtrl(comms.NewMessagesCtrl, "Communication", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "message/:pagesize",
			Name:         "Get Messages",
			Icon:         "fa-box",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/view/:key",
			Name:         "View Message",
			Icon:         "fa-image",
		},
	})

	//Entity
	routes.IdentifyCtrl(entity.NewEntitiesCtrl, "Entity", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/entities/:pagesize",
			Name:         "Get Entities",
			Icon:         "fa-box",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/entity/:key",
			Name:         "Edit Entity",
			Icon:         "fa-image",
		},
	})

	//Folio
	routes.IdentifyCtrl(folio.NewProfileCtrl, "Folio", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/profiles/:pagesize",
			Name:         "Get Entities",
			Icon:         "fa-box",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/profile/:key",
			Name:         "Edit Profile",
			Icon:         "fa-image",
		},
	})

	//Funds
	routes.IdentifyCtrl(funds.NewAccountsCtrl, "Funds", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/accounts/:pagesize",
			Name:         "Get Entities",
			Icon:         "fa-box",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/account/:key",
			Name:         "Edit Account",
			Icon:         "fa-image",
		},
	})

	//Game

	//Gate

	//Logbook

	//Notify

	//Router
	routes.IdentifyCtrl(router.NewMemoryCtrl, "Router", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/memory",
			Name:         "Get Memory",
			Icon:         "fa-microchip",
		},
	})

	//Secure
	routes.IdentifyCtrl(secCtrl.NewUserCtrl, "Secure", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/users/:pagesize",
			Name:         "Get Users",
			Icon:         "fa-group",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/user/:key",
			Name:         "Get User",
			Icon:         "fa-user",
		},
	})

	//Stock
	routes.IdentifyCtrl(stock.NewCarsCtrl, "Stock", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/cars/:pagesize",
			Name:         "Get Cars",
			Icon:         "fa-fleet",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/car/:key",
			Name:         "Get Car",
			Icon:         "fa-car",
		},
	})

	routes.IdentifyCtrl(stock.NewPartsCtrl, "Stock", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/parts/:pagesize",
			Name:         "Get Parts",
			Icon:         "fa-group",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/part/:key",
			Name:         "Get Part",
			Icon:         "fa-user",
		},
	})

	routes.IdentifyCtrl(stock.NewServicesCtrl, "Secure", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/services/:pagesize",
			Name:         "Get Services",
			Icon:         "fa-group",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/service/:key",
			Name:         "Get Service",
			Icon:         "fa-user",
		},
	})

	//Theme

	//Vehicle
	routes.IdentifyCtrl(vehicle.NewVehiclesCtrl, "Vehicle", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/vehicles/:pagesize",
			Name:         "Get Vehicles",
			Icon:         "fa-group",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/vehicle/:key",
			Name:         "Get Vehicle",
			Icon:         "fa-car",
		},
	})

	//Vin
	routes.IdentifyCtrl(vin.NewVINCtrl, "VIN", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/vins/:pagesize",
			Name:         "Get VINs",
			Icon:         "fa-group",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetView",
			Path:         "/vin/:key",
			Name:         "Get VIN",
			Icon:         "fa-car",
		},
	})

	routes.IdentifyCtrl(vin.NewRegionsCtrl, "VIN", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/regions/:pagesize",
			Name:         "Get Regions",
			Icon:         "fa-group",
		},
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "GetEdit",
			Path:         "/region/:key",
			Name:         "Get Region",
			Icon:         "fa-country",
		},
	})

	//XChange
	routes.IdentifyCtrl(xchange.NewCreditCtrl, "XChange", []logic.ControlOption{
		{
			Method:       "GET",
			RequiredRole: roletype.Admin,
			Function:     "Get",
			Path:         "/credit",
			Name:         "Get Credits",
			Icon:         "fa-money",
		},
	})
}
