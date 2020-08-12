package theme

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetStylesheets(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Stylesheets", tmpl, "./views/theme/stylesheets.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStylesheets("A10")

		if err != nil {
			log.Println("Fetch Stylesheets Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchStylesheets(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Stylesheets", tmpl, "./views/theme/stylesheets.html")
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStylesheets(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Stylesheets Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewStylesheets(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Stylesheet View", tmpl, "./views/theme/stylesheetview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStylesheet(key.String())

		if err != nil {
			log.Println("Fetch Stylesheet Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
