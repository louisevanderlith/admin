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
	"github.com/louisevanderlith/admin/controllers/game"
	"github.com/louisevanderlith/admin/controllers/router"
	"github.com/louisevanderlith/admin/controllers/stock"
	"github.com/louisevanderlith/admin/controllers/vehicle"
	"github.com/louisevanderlith/admin/controllers/vin"
	"github.com/louisevanderlith/admin/controllers/xchange"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(e *droxolite.Epoxy) {
	//Home
	homeCtrl := &controllers.Home{}
	homeGroup := routing.NewInterfaceBundle("", roletype.Admin, homeCtrl)

	q := make(map[string]string)
	q["access_token"] = "{access_token}"
	homeGroup.AddRouteWithQueries("Home", "/", "GET", roletype.Admin, q, homeCtrl.Default)
	e.AddGroup(homeGroup)

	//Artifact
	artifactGroup := routing.NewInterfaceBundle("Artifact", roletype.Admin, &artifact.Uploads{})
	e.AddNamedGroup("Artifact.API", artifactGroup)

	//Blog
	blogGroup := routing.NewInterfaceBundle("Blog", roletype.Admin, &blog.Articles{})
	e.AddNamedGroup("Blog.API", blogGroup)

	//Comment
	commentGroup := routing.NewInterfaceBundle("Comment", roletype.Admin, &comment.Messages{})
	e.AddNamedGroup("Comment.API", commentGroup)

	//Comms
	commsGroup := routing.NewInterfaceBundle("Comms", roletype.Admin, &comms.Messages{})
	e.AddNamedGroup("Comms.API", commsGroup)

	//Entity
	entityGroup := routing.NewInterfaceBundle("Entity", roletype.Admin, &entity.Entities{})
	e.AddNamedGroup("Entity.API", entityGroup)

	//Folio
	folioGroup := routing.NewInterfaceBundle("Folio", roletype.Admin, &folio.Profiles{})
	e.AddNamedGroup("Folio.API", folioGroup)

	//subGroup.AddSubGroup(complxGroup)
	//e.AddGroup(subGroup)

	//Funds
	fundsGroup := routing.NewInterfaceBundle("Funds", roletype.Admin, &funds.Accounts{})
	e.AddNamedGroup("Funds.API", fundsGroup)

	//Game
	gameGroup := routing.NewInterfaceBundle("Game", roletype.Admin, &game.Heroes{})
	e.AddNamedGroup("Game.API", gameGroup)

	//Gate

	//Logbook

	//Notify

	//Router
	routerGroup := routing.NewInterfaceBundle("Router", roletype.Admin, &router.Memory{})
	e.AddNamedGroup("Router.API", routerGroup)

	//Secure
	secureGroup := routing.NewInterfaceBundle("Secure", roletype.Admin, &curity.Users{})
	e.AddNamedGroup("Secure.API", secureGroup)

	//Stock
	stockGroup := routing.NewInterfaceBundle("Stock", roletype.Admin, &stock.Stock{}, &stock.Cars{}, &stock.Parts{}, &stock.Services{})
	e.AddNamedGroup("Stock.API", stockGroup)
	//Theme

	//Vehicle
	vehicleGroup := routing.NewInterfaceBundle("Vehicle", roletype.Admin, &vehicle.Vehicles{})
	e.AddNamedGroup("Vehicle.API", vehicleGroup)

	//Vin
	vinGroup := routing.NewInterfaceBundle("VIN", roletype.Admin, &vin.VIN{}, &vin.Regions{})
	e.AddNamedGroup("VIN.API", vinGroup)

	//XChange
	xchangeGroup := routing.NewInterfaceBundle("XChange", roletype.Admin, &xchange.Credits{})
	e.AddNamedGroup("XChange.API", xchangeGroup)
}
