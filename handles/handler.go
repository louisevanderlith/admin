package handles

import (
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/open"
	folio "github.com/louisevanderlith/folio/api"
	"github.com/louisevanderlith/theme/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	CredConfig *clientcredentials.Config
	Endpoints  map[string]string
)

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	authConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID},
	}

	CredConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "theme", "folio"},
	}

	err = api.UpdateTemplate(CredConfig.Client(ctx), endpoints["theme"])

	if err != nil {
		panic(err)
	}

	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	lock := open.NewUILock(authConfig)
	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)

	oidcConfig := &oidc.Config{
		ClientID: clientId,
	}

	v := provider.Verifier(oidcConfig)

	r.HandleFunc("/", open.LoginMiddleware(v, Index(tmpl))).Methods(http.MethodGet)

	r.HandleFunc("/stock", open.LoginMiddleware(v, GetStock(tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/stock/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, ViewStock(tmpl))).Methods(http.MethodGet) //CategorieView

	r.HandleFunc("/cars", open.LoginMiddleware(v, CreateVehicle(tmpl))).Methods(http.MethodGet)                       //Cars
	r.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, ViewVehicle(tmpl))).Methods(http.MethodGet) //CarView
	// //r.HandleFunc("/clothing/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, View(tmpl))).Methods(http.MethodGet)
	// r.HandleFunc("/spares/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, ViewPart(tmpl))).Methods(http.MethodGet)         //PartView
	// r.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, ViewProperty(tmpl))).Methods(http.MethodGet) //PropertyView
	// r.HandleFunc("/utilities/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, ViewService(tmpl))).Methods(http.MethodGet)   //ServicesView
	// r.HandleFunc("/tokens/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, ViewCredits(tmpl))).Methods(http.MethodGet)      //CreditsView

	return r
}

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("e", "/clients", "Clients", nil))
	m.AddItem(menu.NewItem("a", "/orders", "Orders", nil))
	m.AddItem(menu.NewItem("c", "/heroes", "Profiles", nil))

	//TODO: Add categories as children
	m.AddItem(menu.NewItem("b", "/stock", "Stock", nil))

	return m
}

func ThemeContentMod() mix.ModFunc {
	return func(f mix.MixerFactory, r *http.Request) {
		clnt := CredConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		f.SetValue("Folio", content)
	}
}
