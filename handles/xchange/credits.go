package xchange

import (
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

func GetCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Credits", "./views/xchange/credits.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		err := ctx.Serve(http.StatusOK, pge.Page(nil, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Credits", "./views/xchange/credits.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		err := ctx.Serve(http.StatusOK, pge.Page(nil, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "CreditsView", "./views/xchange/creditview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchCredits(key.String())

		if err != nil {
			log.Println("Fetch Credit Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
