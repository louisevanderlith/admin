package handles

import (
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/utility/api"
)

func GetServices(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Services", tmpl, "./views/services.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllServices(clnt, Endpoints["utility"], "A10")

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

func SearchServices(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Services", tmpl, "./views/services.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllServices(clnt, Endpoints["utility"], drx.FindParam(r, "pagesize"))

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

func ViewService(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Service View", tmpl, "./views/serviceview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchService(clnt, Endpoints["utility"], key)

		if err != nil {
			log.Println("Fetch Service Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
