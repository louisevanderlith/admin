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
	credConfig *clientcredentials.Config
	AuthConfig *oauth2.Config
	Endpoints  map[string]string
)

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	AuthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID, "upload-artifact"},
	}

	credConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "theme", "folio"},
	}

	err = api.UpdateTemplate(credConfig.Client(ctx), endpoints["theme"])

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

	lock := open.NewUILock(provider, AuthConfig)
	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)

	r.Handle("/", lock.Middleware(Index(tmpl))).Methods(http.MethodGet)

	r.Handle("/categories", lock.Middleware(GetStockCategories(tmpl))).Methods(http.MethodGet)
	r.Handle("/categories/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewStockCategory(tmpl))).Methods(http.MethodGet)

	r.Handle("/cars", lock.Middleware(GetVehicles(tmpl))).Methods(http.MethodGet)
	r.Handle("/cars/create", lock.Middleware(CreateVehicle(tmpl))).Methods(http.MethodGet)
	r.Handle("/cars/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewVehicle(tmpl))).Methods(http.MethodGet)

	r.Handle("/clothing", lock.Middleware(GetClothing(tmpl))).Methods(http.MethodGet)
	r.Handle("/clothing/create", lock.Middleware(CreateClothing(tmpl))).Methods(http.MethodGet)
	r.Handle("/clothing/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewClothing(tmpl))).Methods(http.MethodGet)

	r.Handle("/spares", lock.Middleware(GetParts(tmpl))).Methods(http.MethodGet)
	r.Handle("/spares/create", lock.Middleware(CreatePart(tmpl))).Methods(http.MethodGet)
	r.Handle("/spares/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewPart(tmpl))).Methods(http.MethodGet)

	r.Handle("/properties", lock.Middleware(GetProperties(tmpl))).Methods(http.MethodGet)
	r.Handle("/properties/create", lock.Middleware(CreateProperty(tmpl))).Methods(http.MethodGet)
	r.Handle("/properties/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewProperty(tmpl))).Methods(http.MethodGet)

	r.Handle("/utilities", lock.Middleware(GetServices(tmpl))).Methods(http.MethodGet)
	r.Handle("/utilities/create", lock.Middleware(CreateService(tmpl))).Methods(http.MethodGet)
	r.Handle("/utilities/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewService(tmpl))).Methods(http.MethodGet)

	//r.HandleFunc("/tokens", lock.Middleware( create(tmpl))).Methods(http.MethodGet)                       //Cars
	r.Handle("/tokens/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewCredits(tmpl))).Methods(http.MethodGet)

	return r
}

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("e", "/clients", "Clients", nil))
	m.AddItem(menu.NewItem("a", "/orders", "Orders", nil))
	m.AddItem(menu.NewItem("c", "/heroes", "Profiles", nil))

	//TODO: Add categories as children
	m.AddItem(menu.NewItem("b", "/categories", "Stock Categories", nil))

	return m
}

func ThemeContentMod() mix.ModFunc {
	return func(f mix.MixerFactory, r *http.Request) {
		clnt := credConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		f.SetValue("Folio", content)
	}
}
