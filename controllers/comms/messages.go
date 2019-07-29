package comms

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type MessagesController struct {
	xontrols.UICtrl
}

func (c *MessagesController) Get() {
	c.Setup("messages", "Messages", true)

	result := []interface{}{}
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Comms.API", "message", "all", "A10")

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *MessagesController) GetView() {
	c.Setup("messageView", "View Message", false)

	key, err := husk.ParseKey(c.FindParam(":key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Comms.API", "message", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
