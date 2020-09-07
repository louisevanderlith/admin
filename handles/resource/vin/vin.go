package vin

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

func GetVIN(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("VIN", tmpl, "./views/vin/vin.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchVINs("A10")

		if err != nil {
			log.Println("Fetch VINs", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchVIN(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("VIN", tmpl, "./views/vin/vin.html")
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchVINs(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch VINs Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewVIN(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("VIN View", tmpl, "./views/vin/vinview.html")
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchVIN(key.String())

		if err != nil {
			log.Println("Fetch VIN Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
