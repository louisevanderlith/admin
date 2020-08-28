package stock

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

func GetCars(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Cars", tmpl, "./views/stock/cars.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchStockCars("A10")

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchCars(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Cars", tmpl, "./views/stock/cars.html")
	return func(w http.ResponseWriter, r *http.Request) {
		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchStockCars(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Car Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewCar(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Car View", tmpl, "./views/stock/carview.html")
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchStockCar(key.String())

		if err != nil {
			log.Println("Fetch Car Erorr", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
