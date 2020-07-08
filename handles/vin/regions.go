package vin

import (
	"fmt"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetRegions(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Regions", "./views/vin/regions.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchRegions("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchRegions(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Regions", "./views/vin/regions.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchRegions(ctx.FindParam("pagesize"))

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

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewRegion(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Regions View", "./views/vin/regionsView.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchRegion(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}
