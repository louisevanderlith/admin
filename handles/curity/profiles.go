package curity

import (
	"fmt"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"html/template"
	"log"
	"net/http"
)

func GetProfiles(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchProfiles("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		result["Next"] = "profiles/B10"
		result["Previous"] = ""
		err = ctx.Serve(http.StatusOK, mix.Page("profiles", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchProfiles(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchProfiles(ctx.FindParam("pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		page, size := ctx.GetPageData()
		result["Next"] = fmt.Sprintf("%c%v", (page+1)+64, size)

		if page != 1 {
			result["Previous"] = fmt.Sprintf("%c%v", (page-1)+64, size)
		}

		err = ctx.Serve(http.StatusOK, mix.Page("profiles", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewProfile(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchProfile(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		//result["Menu"] =

		err = ctx.Serve(http.StatusOK, mix.Page("profilesView", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func CreateProfile(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		err := ctx.Serve(http.StatusOK, mix.Page("profilesCreate", nil, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}


func makeMenu() *bodies.Menu {
	result := bodies.NewMenu()
	result.AddGroup("Home", nil)

	return result
}
