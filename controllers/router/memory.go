package router

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
)

type Memory struct {
	xontrols.UICtrl
}

func (c *Memory) Default() {
	c.Setup("memory", "Memory", true)

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Router.API", "memory")

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
