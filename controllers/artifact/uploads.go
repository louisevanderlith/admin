package artifact

import (
	"fmt"

	"github.com/louisevanderlith/admin/logic"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type UploadsController struct {
	control.UIController
}

func NewUploadsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) logic.PageUI {
	result := &UploadsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (req *UploadsController) Get() {
	req.Setup("uploads", "Uploads", true)

	var result []interface{}
	pagesize := req.Ctx.Input.Param(":pagesize")
	_, err := mango.DoGET(req.GetMyToken(), &result, req.GetInstanceID(), "Artifact.API", "upload", "all", pagesize)

	if err != nil {
		fmt.Println(err)
		req.Serve(nil, err)
		return
	}

	req.Serve(result, nil)
}

func (c *UploadsController) GetView() {
	c.Setup("uploadView", "View Upload", false)

	key, err := husk.ParseKey(c.Ctx.Input.Param(":key"))

	if err != nil {
		c.Serve(nil, err)
	}

	result := make(map[string]interface{})
	_, err = mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Artifact.API", "upload", key.String())

	c.Serve(result, err)
}
