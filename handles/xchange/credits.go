package xchange

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

func GetCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Credits", tmpl, "./views/xchange/credits.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Credits", tmpl, "./views/xchange/credits.html")
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("CreditsView", tmpl, "./views/xchange/creditview.html")
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchCredits(key.String())

		if err != nil {
			log.Println("Fetch Credit Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
