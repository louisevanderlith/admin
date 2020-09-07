package game

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

func GetHeroes(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Heroes", tmpl, "./views/game/heroes.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchHeroes("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchHeroes(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Heroes", tmpl, "./views/game/heroes.html")

	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchHeroes(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewHero(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Hero View", tmpl, "./views/game/heroview.html")

	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchHero(key.String())

		if err != nil {
			log.Println("Serve Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
