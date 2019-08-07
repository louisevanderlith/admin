package xchange

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Credits struct {
	xontrols.UICtrl
}

func (c *Credits) Default() {
	c.Setup("credits", "Credits", true)

	c.Serve(http.StatusNotImplemented, nil, nil)
}

func (c *Credits) Search() {
	c.Setup("credits", "Credits", true)

	c.Serve(http.StatusNotImplemented, nil, nil)
}

func (c *Credits) View() {
	c.Setup("creditView", "View Credit", false)

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "XChange.API", "???", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
