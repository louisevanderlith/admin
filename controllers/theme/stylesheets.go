package theme

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Stylesheets struct {
	xontrols.UICtrl
}

func (c *Stylesheets) Default() {
	c.Setup("stylesheets", "Stylesheets", false)

	var result []interface{}
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Theme.API", "asset", "css")

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Stylesheets) Search() {
	c.Setup("stylesheets", "Stylesheets", false)

	var result []interface{}
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Theme.API", "asset", "css")

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Stylesheets) View() {
	c.Setup("stylesheetView", "View Stylesheet", false)

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Theme.API", "???", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
