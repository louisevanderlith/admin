package stock

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetProperties(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Properties", tmpl, "./views/stock/properties.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockProperties("A10")

		if err != nil {
			log.Println("Fetch Properties Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchProperties(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Properties", tmpl, "./views/stock/properties.html")
	return func(w http.ResponseWriter, r *http.Request) {
		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockProperties(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Properties", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewProperty(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Property View", tmpl, "./views/stock/propertyview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockProperty(key.String())

		if err != nil {
			log.Println("Fetch Property Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
