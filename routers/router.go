package routers

import (
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
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	//Home
	//e.NewInterfaceBundle("", roletype.Admin, &controllers.Home{})

	e.JoinBundle("/artifact", roletype.Admin, mix.Page, &artifact.Uploads{})
	e.JoinBundle("/blog", roletype.Admin, mix.Page, &blog.Articles{})
	e.JoinBundle("/comment", roletype.Admin, mix.Page, &comment.Messages{})
	e.JoinBundle("/comms", roletype.Admin, mix.Page, &comms.Messages{})
	e.JoinBundle("/entity", roletype.Admin, mix.Page, &entity.Entities{})
	e.JoinBundle("/folio", roletype.Admin, mix.Page, &folio.Profiles{})
	e.JoinBundle("/funds", roletype.Admin, mix.Page, &funds.Accounts{})
	e.JoinBundle("/game", roletype.Admin, mix.Page, &game.Heroes{})
	//Gate
	//Logbook
	//Notify
	e.JoinBundle("/router", roletype.Admin, mix.Page, &router.Memory{})
	e.JoinBundle("/secure", roletype.Admin, mix.Page, &curity.Users{})
	e.JoinBundle("/stock", roletype.Admin, mix.Page, &stock.Cars{}, &stock.Parts{}, &stock.Services{})
	//Theme
	e.JoinBundle("/vehicle", roletype.Admin, mix.Page, &vehicle.Vehicles{})
	e.JoinBundle("/vin", roletype.Admin, mix.Page, &vin.VIN{}, &vin.Regions{})

	//XChange -- Needs more development
	//xchangeGroup := routing.NewInterfaceBundle("XChange", roletype.Admin, &xchange.Credits{})
	//e.AddBundle(xchangeGroup)
}
