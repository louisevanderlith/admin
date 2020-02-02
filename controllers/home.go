package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"net/http"
)

func Index(c *gin.Context) {
	result := make(map[string]string)
	result["Comms.API"] = "/comms/messages/A10"
	result["Router.API"] = "/router/memory"
	result["Folio.API"] = "/folio/profiles/A10"

	c.HTML(http.StatusOK, "index.html", droxo.Wrap("Index", result))
}
