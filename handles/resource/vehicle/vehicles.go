package vehicle

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

func GetVehicles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Vehicles", tmpl, "./views/vehicle/vehicles.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchVehicles("A10")

		if err != nil {
			log.Println("Fetch Vehicles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchVehicles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Vehicles", tmpl, "./views/vehicle/vehicles.html")
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchVehicles(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Vehicles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewVehicles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Vehicles View", tmpl, "./views/vehicle/vehicleview.html")
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchVehicle(key.String())

		if err != nil {
			log.Println("Fetch Vehicle Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
