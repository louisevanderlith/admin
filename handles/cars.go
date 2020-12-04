package handles

import (
	"golang.org/x/oauth2"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/vehicle/api"
)

func GetVehicles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllVehicles(clnt, Endpoints["vehicle"], "A10")

		if err != nil {
			log.Println("Fetch Vehicles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Vehicles", "./views/vehicles.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchVehicles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllVehicles(clnt, Endpoints["vehicle"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Vehicles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Vehicles", "./views/vehicles.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateVehicle(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, fact.Create(r, "Vehicle Create", "./views/vehicleview.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewVehicle(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchVehicleInfo(clnt, Endpoints["vehicle"], key)

		if err != nil {
			log.Println("Fetch Vehicle Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Vehicle View", "./views/vehicleview.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
