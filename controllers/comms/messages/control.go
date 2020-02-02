package messages

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/husk"
	"net/http"
)

func Get(c *gin.Context) {
	pagesize := "A10"

	uplURL := fmt.Sprintf("%smessages/%s/", droxo.UriComms, pagesize)
	resp, err := http.Get(uplURL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "comms.html", droxo.Wrap("Comms", result))
}

func Search(c *gin.Context) {
	pagesize := c.Param("pagesize")
	hsh := c.Param("hash")

	uplURL := fmt.Sprintf("%smessages/%s/%s", droxo.UriComms, pagesize, hsh)
	resp, err := http.Get(uplURL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "comms.html", droxo.Wrap("Comms", result))
}

func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uplURL := fmt.Sprintf("%smessage/%s", droxo.UriComms, key)
	resp, err := http.Get(uplURL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "commsView.html", droxo.Wrap("CommsView", result))
}

