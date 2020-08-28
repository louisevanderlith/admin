package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/admin/handles/account/funds"
	"github.com/louisevanderlith/admin/handles/account/game"
	"github.com/louisevanderlith/admin/handles/content"
	"github.com/louisevanderlith/admin/handles/content/artifact"
	"github.com/louisevanderlith/admin/handles/content/blog"
	"github.com/louisevanderlith/admin/handles/profile/comms"
	"github.com/louisevanderlith/admin/handles/resource/stock"
	"github.com/louisevanderlith/admin/handles/resource/vin"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong"
	"html/template"
	"net/http"
)

func SetupRoutes(clnt, scrt, securityUrl, authorityUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, Index(tmpl), map[string]bool{"artifact.uploads.view": true})).Methods(http.MethodGet)
	AddAccountManager(r, clnt, scrt, securityUrl, authorityUrl, tmpl)
	AddContentManager(r, clnt, scrt, securityUrl, authorityUrl, tmpl)
	AddProfileManager(r, clnt, scrt, securityUrl, authorityUrl, tmpl)
	AddResourceManager(r, clnt, scrt, securityUrl, authorityUrl, tmpl)
	return r
}

func AddAccountManager(r *mux.Router, clnt, scrt, securityUrl, authorityUrl string, tmpl *template.Template) {
	managr := r.PathPrefix("/accountmanage").Subrouter()

	gme := managr.PathPrefix("/game").Subrouter()
	gme.HandleFunc("/heroes", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, game.GetHeroes(tmpl), map[string]bool{"game.heroes.search": true})).Methods(http.MethodGet)

	fund := managr.PathPrefix("/fund").Subrouter()
	fund.HandleFunc("/accounts", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, funds.GetAccounts(tmpl), map[string]bool{"funds.account.search": true})).Methods(http.MethodGet)
}

func AddContentManager(r *mux.Router, clnt, scrt, securityUrl, authorityUrl string, tmpl *template.Template) {
	managr := r.PathPrefix("/contentmanage").Subrouter()
	managr.HandleFunc("", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, content.Management(tmpl), map[string]bool{"artifact.uploads.view": true, "blog.articles.view": true})).Methods(http.MethodGet)
	artfct := managr.PathPrefix("/artifact").Subrouter()

	artfct.HandleFunc("/uploads", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.GetUploads(tmpl), map[string]bool{"artifact.uploads.view": true})).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.SearchUploads(tmpl), map[string]bool{"artifact.uploads.search": true})).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.SearchUploads(tmpl), map[string]bool{"artifact.uploads.search": true})).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.ViewUpload(tmpl), map[string]bool{"artifact.uploads.view": true})).Methods(http.MethodGet)

	blg := managr.PathPrefix("/blog").Subrouter()
	blg.HandleFunc("/articles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.GetArticles(tmpl), map[string]bool{"blog.articles.view": true})).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.SearchArticles(tmpl), map[string]bool{"blog.articles.search": true})).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.SearchArticles(tmpl), map[string]bool{"blog.articles.search": true})).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.ViewArticle(tmpl), map[string]bool{"blog.articles.view": true})).Methods(http.MethodGet)

}

func AddProfileManager(r *mux.Router, clnt, scrt, securityUrl, authorityUrl string, tmpl *template.Template) {
	managr := r.PathPrefix("/profilemanage").Subrouter()

	cmms := managr.PathPrefix("/comms").Subrouter()
	cmms.HandleFunc("/messages", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.GetMessages(tmpl), map[string]bool{"comms.messages.search": true})).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.SearchMessages(tmpl), map[string]bool{"comms.messages.search": true})).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.SearchMessages(tmpl), map[string]bool{"comms.messages.search": true})).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.ViewMessage(tmpl), map[string]bool{"comms.messages.view": true})).Methods(http.MethodGet)

}

func AddResourceManager(r *mux.Router, clnt, scrt, securityUrl, authorityUrl string, tmpl *template.Template) {
	managr := r.PathPrefix("/resourcemanage").Subrouter()

	stck := managr.PathPrefix("/stock").Subrouter()
	stck.HandleFunc("/cars", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.GetCars(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchCars(tmpl), map[string]bool{"stock.cars.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchCars(tmpl), map[string]bool{"stock.cars.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.ViewCar(tmpl), map[string]bool{"stock.cars.view": true})).Methods(http.MethodGet)

	stck.HandleFunc("/parts", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.GetParts(tmpl), map[string]bool{"stock.parts.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchParts(tmpl), map[string]bool{"stock.parts.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchParts(tmpl), map[string]bool{"stock.parts.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.ViewPart(tmpl), map[string]bool{"stock.parts.view": true})).Methods(http.MethodGet)

	stck.HandleFunc("/properties", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.GetProperties(tmpl), map[string]bool{"stock.properties.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchProperties(tmpl), map[string]bool{"stock.properties.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchProperties(tmpl), map[string]bool{"stock.properties.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.ViewProperty(tmpl), map[string]bool{"stock.properties.view": true})).Methods(http.MethodGet)

	stck.HandleFunc("/services", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.GetServices(tmpl), map[string]bool{"stock.services.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/services/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchServices(tmpl), map[string]bool{"stock.services.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.SearchServices(tmpl), map[string]bool{"stock.services.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, stock.ViewProperty(tmpl), map[string]bool{"stock.services.view": true})).Methods(http.MethodGet)

	vins := managr.PathPrefix("/vin").Subrouter()
	vins.HandleFunc("/regions", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.GetRegions(tmpl), map[string]bool{"vin.region.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchRegions(tmpl), map[string]bool{"vin.region.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchRegions(tmpl), map[string]bool{"vin.region.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.ViewRegion(tmpl), map[string]bool{"vin.region.view": true})).Methods(http.MethodGet)

	vins.HandleFunc("/numbers", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.GetVIN(tmpl), map[string]bool{"vin.admin.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchVIN(tmpl), map[string]bool{"vin.admin.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchVIN(tmpl), map[string]bool{"vin.admin.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.ViewVIN(tmpl), map[string]bool{"vin.admin.view": true})).Methods(http.MethodGet)
}
