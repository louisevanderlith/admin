package vin

import (
	"fmt"
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetRegions(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Regions", tmpl, "./views/vin/regions.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchRegions("A10")

		if err != nil {
			log.Println("Fetch Regions Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchRegions(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Regions", tmpl, "./views/vin/regions.html")
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchRegions(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Regions", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		page, size := drx.GetPageData(r)
		result["Next"] = fmt.Sprintf("%c%v", (page+1)+64, size)

		if page != 1 {
			result["Previous"] = fmt.Sprintf("%c%v", (page-1)+64, size)
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewRegion(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Region View", tmpl, "./views/vin/regionview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchRegion(key.String())

		if err != nil {
			log.Println("Fetch Region Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
