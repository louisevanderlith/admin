package stock

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetServices(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Services", "./views/stock/services.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchStockCars("A10")

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchServices(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Services", "./views/stock/services.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchStockCars(ctx.FindParam("pagesize"))

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewService(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Service View", "./views/stock/serviceview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)

		result, err := src.FetchStockService(key.String())

		if err != nil {
			log.Println("Fetch Service Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
