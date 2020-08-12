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

func GetTemplates(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Templates", tmpl, "./views/theme/templates.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchTemplates("A10")

		if err != nil {
			log.Println("Fetch Templates Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchTemplates(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Templates", tmpl, "./views/theme/templates.html")
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchTemplates(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Templates Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewTemplates(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Template View", tmpl, "./views/theme/templateview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchTemplate(key.String())

		if err != nil {
			log.Println("Fetch Template Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
