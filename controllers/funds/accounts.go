package funds

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type Accounts struct {
	xontrols.UICtrl
}

func (c *Accounts) Default() {
	c.Setup("accounts", "Accounts", true)

	var result []interface{}
	pagesize := "A10"

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Funds.API", "account", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Accounts) Search() {
	c.Setup("accounts", "Accounts", true)

	var result []interface{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Funds.API", "account", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *Accounts) View() {
	c.Setup("accountEdit", "Edit Account", true)
	c.EnableSave()

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Funds.API", "account", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
