package handles

import (
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/mix"
)

func GetCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Credits", tmpl, "./views/credits.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Credits", tmpl, "./views/credits.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CreateCredit(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("CreditsView", tmpl, "./views/creditview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("CreditsView", tmpl, "./views/creditview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
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
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
