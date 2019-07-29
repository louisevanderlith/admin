package artifact

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type UploadsController struct {
	xontrols.UICtrl
}

func (req *UploadsController) Get() {
	req.Setup("uploads", "Uploads", true)

	var result []interface{}
	pagesize := req.FindParam("pagesize")
	code, err := droxolite.DoGET(req.GetMyToken(), &result, req.Settings.InstanceID, "Artifact.API", "upload", "all", pagesize)

	if err != nil {
		log.Println(err)
		req.Serve(code, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}

func (c *UploadsController) GetView() {
	c.Setup("uploadView", "View Upload", false)

	key, err := husk.ParseKey(c.FindParam("key"))

	if err != nil {
		c.Serve(http.StatusBadRequest, err, nil)
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Artifact.API", "upload", key.String())

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}
