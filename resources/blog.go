package resources

import (
	"log"
)

func (src *Source) FetchArticle(key string) (interface{}, error) {
	return src.get("blog", "articles", key)
}

func (src *Source) FetchArticles(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("blog", "articles", pagesize)

	if err != nil {
		return nil, err
	}

	obj, ok := res.(map[string]interface{})

	if !ok {
		log.Println("Cast Error", res)
		return make(map[string]interface{}), nil
	}

	return obj, nil
}
