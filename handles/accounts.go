package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/funds/api"
	"github.com/louisevanderlith/husk/keys"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GetAccounts(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllAccounts(clnt, Endpoints["funds"], "A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Accounts", "./views/funds/accounts.html", b))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchAccounts(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllAccounts(clnt, Endpoints["funds"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Accounts", "./views/funds/accounts.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewAccounts(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAccount(clnt, Endpoints["funds"], key)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		b := mix.NewDataBag(data)
		err = mix.Write(w, fact.Create(r, "Account View", "./views/accountview.html", b))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
