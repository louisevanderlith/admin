package comms

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type MessagesController struct {
	control.UIController
}

func NewMessagesCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) beego.ControllerInterface {
	result := &MessagesController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *MessagesController) Get() {
	c.Setup("messages", "Messages", true)

	result := []interface{}{}
	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Comms.API", "message", "all", "A10")

	c.Serve(result, err)
}

func (c *MessagesController) GetView() {
	c.Setup("messageView", "View Message", false)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Comms.API", "message", key.String())

	if err != nil {
		log.Println(err)
	}

	c.Serve(result, err)
}
