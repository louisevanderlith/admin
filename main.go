package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/admin/controllers"
	"github.com/louisevanderlith/droxo"
	"os"
)

func main() {
	prof := os.Getenv("PROFILE")

	if len(prof) == 0 {
		panic("invalid profile")
	}

	host := os.Getenv("HOST")
	authority := fmt.Sprintf("https://oauth2.%s", host)

	droxo.AssignOperator(prof, host)
	droxo.DefineClient("admin", "adminsecret", host, authority)
	//Download latest Theme
	err := droxo.UpdateTheme("http://theme:8093")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	tmpl, err := droxo.LoadTemplates("./views")

	if err != nil {
		panic(err)
	}

	r.HTMLRender = tmpl

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))
	r.Use(droxo.AuthorizeClient())
	r.GET("/", controllers.Index)

	/*
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
	*/

	r.POST("/oauth2", droxo.AuthCallback)

	err = r.Run(":8088")

	if err != nil {
		panic(err)
	}
}
