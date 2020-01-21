package comment

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func Get(c *gin.Context) {
	pagesize := "A10"

	uplURL := fmt.Sprintf("%smessages/%s/", droxo.UriComment, pagesize)
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

	c.HTML(http.StatusOK, "comments.html", droxo.Wrap("Comments", result))
}

func Search(c *gin.Context) {
	pagesize := c.Param("pagesize")
	hsh := c.Param("hash")

	uplURL := fmt.Sprintf("%smessages/%s/%s", droxo.UriComment, pagesize, hsh)
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

	c.HTML(http.StatusOK, "comments.html", droxo.Wrap("Comments", result))
}

func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uplURL := fmt.Sprintf("%smessage/%s", droxo.UriComment, key)
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

	c.HTML(http.StatusOK, "commentView.html", droxo.Wrap("CommentView", result))
}
