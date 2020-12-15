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
		Scopes:       []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess, "upload-artifact"},
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

	lock := open.NewHybridLock(provider, credConfig, AuthConfig)

	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)
	r.HandleFunc("/logout", lock.Logout).Methods(http.MethodGet)
	r.HandleFunc("/refresh", lock.Refresh).Methods(http.MethodGet)

	fact := mix.NewPageFactory(tmpl)
	fact.AddMenu(FullMenu())
	fact.AddModifier(mix.EndpointMod(Endpoints))
	fact.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	fact.AddModifier(ThemeContentMod())

	r.Handle("/", lock.Protect(lock.Lock(Index(fact)))).Methods(http.MethodGet)

	rcat := r.PathPrefix("/categories").Subrouter()
	rcat.Handle("", GetStockCategories(fact)).Methods(http.MethodGet)
	rcat.Handle("/{pagesize:[A-Z][0-9]+}", SearchStockCategories(fact)).Methods(http.MethodGet)
	rcat.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchStockCategories(fact)).Methods(http.MethodGet)
	rcat.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewStockCategory(fact)).Methods(http.MethodGet)
	rcat.Use(lock.Protect)
	rcat.Use(lock.Lock)

	rcar := r.PathPrefix("/cars").Subrouter()
	rcar.Handle("", GetVehicles(fact)).Methods(http.MethodGet)
	rcar.Handle("/create", CreateVehicle(fact)).Methods(http.MethodGet)
	rcar.Handle("/{pagesize:[A-Z][0-9]+}", SearchVehicles(fact)).Methods(http.MethodGet)
	rcar.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchVehicles(fact)).Methods(http.MethodGet)
	rcar.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewVehicle(fact)).Methods(http.MethodGet)
	rcar.Use(lock.Protect)
	rcar.Use(lock.Lock)

	rcloth := r.PathPrefix("/clothing").Subrouter()
	rcloth.Handle("", GetClothing(fact)).Methods(http.MethodGet)
	rcloth.Handle("/create", CreateClothing(fact)).Methods(http.MethodGet)
	rcloth.Handle("/{pagesize:[A-Z][0-9]+}", SearchClothing(fact)).Methods(http.MethodGet)
	rcloth.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchClothing(fact)).Methods(http.MethodGet)
	rcloth.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewClothing(fact)).Methods(http.MethodGet)
	rcloth.Use(lock.Protect)
	rcloth.Use(lock.Lock)

	rpart := r.PathPrefix("/spares").Subrouter()
	rpart.Handle("", GetParts(fact)).Methods(http.MethodGet)
	rpart.Handle("/create", CreatePart(fact)).Methods(http.MethodGet)
	rpart.Handle("/{pagesize:[A-Z][0-9]+}", SearchParts(fact)).Methods(http.MethodGet)
	rpart.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchParts(fact)).Methods(http.MethodGet)
	rpart.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewPart(fact)).Methods(http.MethodGet)
	rpart.Use(lock.Protect)
	rpart.Use(lock.Lock)

	rhouse := r.PathPrefix("/properties").Subrouter()
	rhouse.Handle("", GetProperties(fact)).Methods(http.MethodGet)
	rhouse.Handle("/create", CreateProperty(fact)).Methods(http.MethodGet)
	rhouse.Handle("/{pagesize:[A-Z][0-9]+}", SearchProperties(fact)).Methods(http.MethodGet)
	rhouse.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchProperties(fact)).Methods(http.MethodGet)
	rhouse.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewProperty(fact)).Methods(http.MethodGet)
	rhouse.Use(lock.Protect)
	rhouse.Use(lock.Lock)

	rutil := r.PathPrefix("/utilities").Subrouter()
	rutil.Handle("", GetServices(fact)).Methods(http.MethodGet)
	rutil.Handle("/create", CreateService(fact)).Methods(http.MethodGet)
	rutil.Handle("/{pagesize:[A-Z][0-9]+}", SearchServices(fact)).Methods(http.MethodGet)
	rutil.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchServices(fact)).Methods(http.MethodGet)
	rutil.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewService(fact)).Methods(http.MethodGet)
	rutil.Use(lock.Protect)
	rutil.Use(lock.Lock)

	//r.HandleFunc("/tokens", lock.Middleware( create(tmpl))).Methods(http.MethodGet)                       //Cars
	r.Handle("/tokens/{key:[0-9]+\\x60[0-9]+}", lock.Protect(lock.Lock(ViewCredits(fact)))).Methods(http.MethodGet)

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
	return func(b mix.Bag, r *http.Request) {
		clnt := credConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		b.SetValue("Folio", content)
	}
}
