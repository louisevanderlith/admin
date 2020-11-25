package handles

import (
	"github.com/louisevanderlith/wear/core"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/wear/api"
)

func GetClothing(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Clothing", tmpl, "./views/clothing.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchAllClothing(clnt, Endpoints["wear"], "A10")

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

func SearchClothing(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Clothing", tmpl, "./views/clothing.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchAllClothing(clnt, Endpoints["wear"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Clothing Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateClothing(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Clothing Create", tmpl, "./views/clothingview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		brands, err := api.FetchAllBrands(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}
		pge.SetValue("Brands", brands)

		types, err := api.FetchAllTypes(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println("Fetch Types Error", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		pge.SetValue("Types", types)
		err = mix.Write(w, pge.Create(r, core.Clothing{}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewClothing(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Clothing View", tmpl, "./views/clothingview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		brands, err := api.FetchAllBrands(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}
		pge.SetValue("Brands", brands)

		types, err := api.FetchAllTypes(clnt, Endpoints["wear"], "A10")

		if err != nil {
			log.Println("Fetch Types Error", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		pge.SetValue("Types", types)
		result, err := api.FetchClothing(clnt, Endpoints["wear"], key)

		if err != nil {
			log.Println("Fetch Clothing Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
