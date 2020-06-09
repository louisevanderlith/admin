package vehicle

import (
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetVehicles(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchVehicles("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, mix.Page("vehicles", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchVehicles(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchVehicles(ctx.FindParam("pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, mix.Page("vehicles", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewVehicles(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchVehicle(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, mix.Page("vehiclesView", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}
