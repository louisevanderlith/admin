package uploads

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/husk"
)

func Get(c *gin.Context) {
	pagesize := "A10"

	uplURL := fmt.Sprintf("%suploads/%s/", droxo.UriArtifact, pagesize)
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

	c.HTML(http.StatusOK, "uploads.html", droxo.Wrap("Uploads", result))
}

func Search(c *gin.Context) {
	pagesize := c.Param("pagesize")
	hash := c.Param("hash")
	uplURL := fmt.Sprintf("%suploads/%s/%s", droxo.UriArtifact, pagesize, hash)
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

	c.HTML(http.StatusOK, "uploads.html", droxo.Wrap("Uploads", result))
}

func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uplURL := fmt.Sprintf("%supload/%s/", droxo.UriArtifact, key)
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

	c.HTML(http.StatusOK, "uploadView.html", droxo.Wrap("UploadView", result))
}
