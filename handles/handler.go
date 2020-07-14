package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/admin/handles/artifact"
	"github.com/louisevanderlith/admin/handles/blog"
	"github.com/louisevanderlith/admin/handles/comment"
	"github.com/louisevanderlith/admin/handles/comms"
	"github.com/louisevanderlith/admin/handles/curity"
	"github.com/louisevanderlith/admin/handles/entity"
	"github.com/louisevanderlith/admin/handles/funds"
	"github.com/louisevanderlith/admin/handles/game"
	"github.com/louisevanderlith/admin/handles/vin"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/kong"
	"net/http"
)

func SetupRoutes(clnt, scrt, secureUrl, authUrl string) http.Handler {
	tmpl, err := droxolite.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, Index(tmpl), "artifact.uploads.view")).Methods(http.MethodGet)

	artfct := r.PathPrefix("/artifact").Subrouter()
	artfct.HandleFunc("/uploads", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, artifact.GetUploads(tmpl), "artifact.uploads.view")).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, artifact.SearchUploads(tmpl), "artifact.uploads.search")).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, artifact.SearchUploads(tmpl), "artifact.uploads.search")).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, artifact.ViewUpload(tmpl), "artifact.uploads.view")).Methods(http.MethodGet)

	blg := r.PathPrefix("/blog").Subrouter()
	blg.HandleFunc("/articles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, blog.GetArticles(tmpl), "blog.articles.view")).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, blog.SearchArticles(tmpl), "blog.articles.search")).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, blog.SearchArticles(tmpl), "blog.articles.search")).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, blog.ViewArticle(tmpl), "blog.articles.view")).Methods(http.MethodGet)

	cmmnt := r.PathPrefix("/comment").Subrouter()
	cmmnt.HandleFunc("/messages", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, comment.GetMessages(tmpl), "artifact.uploads.view")).Methods(http.MethodGet)

	gme := r.PathPrefix("/game").Subrouter()
	gme.HandleFunc("/heroes", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, game.GetHeroes(tmpl), "game.heroes.search")).Methods(http.MethodGet)

	fund := r.PathPrefix("/fund").Subrouter()
	fund.HandleFunc("/accounts", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, funds.GetAccounts(tmpl), "funds.account.search")).Methods(http.MethodGet)

	cmms := r.PathPrefix("/comms").Subrouter()
	cmms.HandleFunc("/messages", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, comms.GetMessages(tmpl), "comms.messages.search")).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, comms.SearchMessages(tmpl), "comms.messages.search")).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, comms.SearchMessages(tmpl), "comms.messages.search")).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, comms.ViewMessage(tmpl), "comms.messages.view")).Methods(http.MethodGet)

	vins := r.PathPrefix("/vin").Subrouter()
	vins.HandleFunc("/regions", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.GetRegions(tmpl), "vin.region.search")).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.SearchRegions(tmpl), "vin.region.search")).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.SearchRegions(tmpl), "vin.region.search")).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.ViewRegion(tmpl), "vin.region.view")).Methods(http.MethodGet)

	vins.HandleFunc("/numbers", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.GetVIN(tmpl), "vin.admin.search")).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.SearchVIN(tmpl), "vin.admin.search")).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.SearchVIN(tmpl), "vin.admin.search")).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, vin.ViewVIN(tmpl), "vin.admin.view")).Methods(http.MethodGet)

	ent := r.PathPrefix("/entity").Subrouter()
	ent.HandleFunc("/entities", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.GetEnitites(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	ent.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	ent.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	ent.HandleFunc("/entities/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.ViewEntity(tmpl), "secure.profile.view")).Methods(http.MethodGet)

	crty := r.PathPrefix("/curity").Subrouter()
	crty.HandleFunc("/profiles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.GetProfiles(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchProfiles(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchProfiles(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.ViewProfile(tmpl), "secure.profile.view")).Methods(http.MethodGet)

	crty.HandleFunc("/resources", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.GetResource(tmpl), "secure.resource.search")).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchResource(tmpl), "secure.resource.search")).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchResource(tmpl), "secure.resource.search")).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.ViewResource(tmpl), "secure.resource.view")).Methods(http.MethodGet)

	crty.HandleFunc("/users", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.GetUsers(tmpl), "secure.user.search")).Methods(http.MethodGet)
	crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchUsers(tmpl), "secure.user.search")).Methods(http.MethodGet)
	crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchUsers(tmpl), "secure.user.search")).Methods(http.MethodGet)
	crty.HandleFunc("/users/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.ViewUser(tmpl), "secure.user.view")).Methods(http.MethodGet)

	stck := r.PathPrefix("/stock").Subrouter()
	stck.HandleFunc("/cars", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.GetEnitites(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.cars.search")).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.cars.search")).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.ViewEntity(tmpl), "stock.cars.view")).Methods(http.MethodGet)

	stck.HandleFunc("/parts", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.GetEnitites(tmpl), "stock.parts.search")).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.parts.search")).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.parts.search")).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.ViewEntity(tmpl), "stock.parts.view")).Methods(http.MethodGet)

	stck.HandleFunc("/properties", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.GetEnitites(tmpl), "stock.properties.search")).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.properties.search")).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.properties.search")).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.ViewEntity(tmpl), "stock.properties.view")).Methods(http.MethodGet)

	stck.HandleFunc("/services", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.GetEnitites(tmpl), "stock.services.search")).Methods(http.MethodGet)
	stck.HandleFunc("/services/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.services.search")).Methods(http.MethodGet)
	stck.HandleFunc("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.SearchEntities(tmpl), "stock.services.search")).Methods(http.MethodGet)
	stck.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, entity.ViewEntity(tmpl), "stock.services.view")).Methods(http.MethodGet)

	return r
}
