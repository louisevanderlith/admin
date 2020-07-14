package theme

import (
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

func GetTemplates(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Templates", "./views/theme/templates.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchTemplates("A10")

		if err != nil {
			log.Println("Fetch Templates Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchTemplates(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Templates", "./views/theme/templates.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchTemplates(ctx.FindParam("pagesize"))

		if err != nil {
			log.Println("Fetch Templates Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewTemplates(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Template View", "./views/theme/templateview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchTemplate(key.String())

		if err != nil {
			log.Println("Fetch Template Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
