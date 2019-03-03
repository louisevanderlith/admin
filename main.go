package main

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/admin/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
)

func main() {
	mode := os.Getenv("RUNMODE")

	// Register with router
	name := beego.BConfig.AppName
	srv := mango.NewService(mode, name, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
