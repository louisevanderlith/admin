package entity

import (
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetEnitites(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchEntities("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, mix.Page("entities", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchEntities(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)

		result, err := src.FetchEntities(ctx.FindParam("pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, mix.Page("entities", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewEntity(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchEntity(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, mix.Page("entitiesView", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}
