package articles

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

	uplURL := fmt.Sprintf("%sarticle/%s/", droxo.UriBlog, pagesize)
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

	c.HTML(http.StatusOK, "articles", droxo.Wrap("Articles", result))
}

func Search(c *gin.Context) {
	pagesize := c.Param("pagesize")
	hash := c.Param("hash")
	uplURL := fmt.Sprintf("%sarticles/%s/%s", droxo.UriBlog, pagesize, hash)
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

	c.HTML(http.StatusOK, "articles.html", droxo.Wrap("Articles", result))
}

func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uplURL := fmt.Sprintf("%sarticle/%s", droxo.UriBlog, key)
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

	c.HTML(http.StatusOK, "articleView.html", droxo.Wrap("ArticleView", result))
}

func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "articleCreate.html", droxo.Wrap("ArticleCreate", nil))
}
