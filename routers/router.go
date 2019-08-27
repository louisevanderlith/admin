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
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(e resins.Epoxi) {
	//Home
	homeGroup := routing.NewInterfaceBundle("", roletype.Admin, &controllers.Home{})
	e.AddGroup(homeGroup)

	//Artifact
	artifactGroup := routing.NewInterfaceBundle("Artifact", roletype.Admin, &artifact.Uploads{})
	e.AddGroup(artifactGroup)

	//Blog
	blogGroup := routing.NewInterfaceBundle("Blog", roletype.Admin, &blog.Articles{})
	e.AddGroup(blogGroup)

	//Comment
	commentGroup := routing.NewInterfaceBundle("Comment", roletype.Admin, &comment.Messages{})
	e.AddGroup(commentGroup)

	//Comms
	commsGroup := routing.NewInterfaceBundle("Comms", roletype.Admin, &comms.Messages{})
	e.AddGroup(commsGroup)

	//Entity
	entityGroup := routing.NewInterfaceBundle("Entity", roletype.Admin, &entity.Entities{})
	e.AddGroup(entityGroup)

	//Folio
	folioGroup := routing.NewInterfaceBundle("Folio", roletype.Admin, &folio.Profiles{})
	e.AddGroup(folioGroup)

	//Funds
	fundsGroup := routing.NewInterfaceBundle("Funds", roletype.Admin, &funds.Accounts{})
	e.AddGroup(fundsGroup)

	//Game
	gameGroup := routing.NewInterfaceBundle("Game", roletype.Admin, &game.Heroes{})
	e.AddGroup(gameGroup)

	//Gate

	//Logbook

	//Notify

	//Router
	routerGroup := routing.NewInterfaceBundle("Router", roletype.Admin, &router.Memory{})
	e.AddGroup(routerGroup)

	//Secure
	secureGroup := routing.NewInterfaceBundle("Secure", roletype.Admin, &curity.Users{})
	e.AddGroup(secureGroup)

	//Stock
	stockGroup := routing.NewInterfaceBundle("Stock", roletype.Admin, &stock.Cars{}, &stock.Parts{}, &stock.Services{})
	e.AddGroup(stockGroup)
	//Theme

	//Vehicle
	vehicleGroup := routing.NewInterfaceBundle("Vehicle", roletype.Admin, &vehicle.Vehicles{})
	e.AddGroup(vehicleGroup)

	//Vin
	vinGroup := routing.NewInterfaceBundle("VIN", roletype.Admin, &vin.VIN{}, &vin.Regions{})
	e.AddGroup(vinGroup)

	//XChange -- Needs more development
	//xchangeGroup := routing.NewInterfaceBundle("XChange", roletype.Admin, &xchange.Credits{})
	//e.AddGroup(xchangeGroup)
}
