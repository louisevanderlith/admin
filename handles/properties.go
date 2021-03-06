package handles

import (
	"golang.org/x/oauth2"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/house/api"
	"github.com/louisevanderlith/husk/keys"
)

func GetProperties(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllProperties(clnt, Endpoints["house"], "A10")

		if err != nil {
			log.Println("Fetch Properties Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Properties", "./views/properties.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchProperties(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllProperties(clnt, Endpoints["house"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Properties", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Properties", "./views/properties.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateProperty(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, fact.Create(r, "Property Create", "./views/propertyview.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewProperty(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchProperty(clnt, Endpoints["house"], key)

		if err != nil {
			log.Println("Fetch Property Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Property View", "./views/propertyview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
