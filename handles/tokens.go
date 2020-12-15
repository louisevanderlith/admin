package handles

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/mix"
)

func GetCredits(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Credits", "./views/credits.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchCredits(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Credits", "./views/credits.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateCredit(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, fact.Create(r, "Credit Create", "./views/creditview.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewCredits(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		*/
		/*tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchCredits(key.String())

		if err != nil {
			log.Println("Fetch Credit Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		*/
		err := mix.Write(w, fact.Create(r, "Credit View", "./views/creditview.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
