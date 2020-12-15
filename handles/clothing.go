package handles

import (
	"github.com/louisevanderlith/wear/core"
	"golang.org/x/oauth2"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/wear/api"
)

func GetClothing(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllClothing(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Clothing", "./views/clothing.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchClothing(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllClothing(clnt, Endpoints["wear"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Clothing Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Clothing", "./views/clothing.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateClothing(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		brands, err := api.FetchAllBrands(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		b := mix.NewDataBag(core.Clothing{})
		b.SetValue("Brands", brands)

		types, err := api.FetchAllTypes(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println("Fetch Types Error", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		b.SetValue("Types", types)
		err = mix.Write(w, fact.Create(r, "Clothing Create", "./views/clothingview.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewClothing(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)

		data, err := api.FetchClothing(clnt, Endpoints["wear"], key)

		if err != nil {
			log.Println("Fetch Clothing Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)

		brands, err := api.FetchAllBrands(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		b.SetValue("Brands", brands)

		types, err := api.FetchAllTypes(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println("Fetch Types Error", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		b.SetValue("Types", types)

		err = mix.Write(w, fact.Create(r, "Clothing View", "./views/clothingview.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
