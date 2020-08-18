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
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong"
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

	artfct := r.PathPrefix("/artifact").Subrouter()
	artfct.HandleFunc("/uploads", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.GetUploads(tmpl), map[string]bool{"artifact.uploads.view": true})).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.SearchUploads(tmpl), map[string]bool{"artifact.uploads.search": true})).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.SearchUploads(tmpl), map[string]bool{"artifact.uploads.search": true})).Methods(http.MethodGet)
	artfct.HandleFunc("/uploads/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, artifact.ViewUpload(tmpl), map[string]bool{"artifact.uploads.view": true})).Methods(http.MethodGet)

	blg := r.PathPrefix("/blog").Subrouter()
	blg.HandleFunc("/articles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.GetArticles(tmpl), map[string]bool{"blog.articles.view": true})).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.SearchArticles(tmpl), map[string]bool{"blog.articles.search": true})).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.SearchArticles(tmpl), map[string]bool{"blog.articles.search": true})).Methods(http.MethodGet)
	blg.HandleFunc("/articles/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, blog.ViewArticle(tmpl), map[string]bool{"blog.articles.view": true})).Methods(http.MethodGet)

	cmmnt := r.PathPrefix("/comment").Subrouter()
	cmmnt.HandleFunc("/messages", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comment.GetMessages(tmpl), map[string]bool{"artifact.uploads.view": true})).Methods(http.MethodGet)

	gme := r.PathPrefix("/game").Subrouter()
	gme.HandleFunc("/heroes", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, game.GetHeroes(tmpl), map[string]bool{"game.heroes.search": true})).Methods(http.MethodGet)

	fund := r.PathPrefix("/fund").Subrouter()
	fund.HandleFunc("/accounts", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, funds.GetAccounts(tmpl), map[string]bool{"funds.account.search": true})).Methods(http.MethodGet)

	cmms := r.PathPrefix("/comms").Subrouter()
	cmms.HandleFunc("/messages", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.GetMessages(tmpl), map[string]bool{"comms.messages.search": true})).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.SearchMessages(tmpl), map[string]bool{"comms.messages.search": true})).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.SearchMessages(tmpl), map[string]bool{"comms.messages.search": true})).Methods(http.MethodGet)
	cmms.HandleFunc("/messages/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, comms.ViewMessage(tmpl), map[string]bool{"comms.messages.view": true})).Methods(http.MethodGet)

	vins := r.PathPrefix("/vin").Subrouter()
	vins.HandleFunc("/regions", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.GetRegions(tmpl), map[string]bool{"vin.region.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchRegions(tmpl), map[string]bool{"vin.region.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchRegions(tmpl), map[string]bool{"vin.region.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/regions/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.ViewRegion(tmpl), map[string]bool{"vin.region.view": true})).Methods(http.MethodGet)

	vins.HandleFunc("/numbers", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.GetVIN(tmpl), map[string]bool{"vin.admin.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchVIN(tmpl), map[string]bool{"vin.admin.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.SearchVIN(tmpl), map[string]bool{"vin.admin.search": true})).Methods(http.MethodGet)
	vins.HandleFunc("/numbers/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, vin.ViewVIN(tmpl), map[string]bool{"vin.admin.view": true})).Methods(http.MethodGet)

	ent := r.PathPrefix("/entity").Subrouter()
	ent.HandleFunc("/entities", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.GetEnitites(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)
	ent.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)
	ent.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)
	ent.HandleFunc("/entities/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.ViewEntity(tmpl), map[string]bool{"entity.info.view": true})).Methods(http.MethodGet)

	crty := r.PathPrefix("/curity").Subrouter()
	crty.HandleFunc("/profiles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.GetProfiles(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.SearchProfiles(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.SearchProfiles(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.ViewProfile(tmpl), map[string]bool{"secure.profile.view": true})).Methods(http.MethodGet)

	crty.HandleFunc("/resources", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.GetResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.SearchResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.SearchResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.ViewResource(tmpl), map[string]bool{"secure.resource.view": true})).Methods(http.MethodGet)

	crty.HandleFunc("/users", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.GetUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.SearchUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.SearchUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	crty.HandleFunc("/users/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, curity.ViewUser(tmpl), map[string]bool{"secure.user.view": true})).Methods(http.MethodGet)

	stck := r.PathPrefix("/stock").Subrouter()
	stck.HandleFunc("/cars", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.GetEnitites(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.cars.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.cars.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.ViewEntity(tmpl), map[string]bool{"stock.cars.view": true})).Methods(http.MethodGet)

	stck.HandleFunc("/parts", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.GetEnitites(tmpl), map[string]bool{"stock.parts.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.parts.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.parts.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.ViewEntity(tmpl), map[string]bool{"stock.parts.view": true})).Methods(http.MethodGet)

	stck.HandleFunc("/properties", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.GetEnitites(tmpl), map[string]bool{"stock.properties.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.properties.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.properties.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.ViewEntity(tmpl), map[string]bool{"stock.properties.view": true})).Methods(http.MethodGet)

	stck.HandleFunc("/services", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.GetEnitites(tmpl), map[string]bool{"stock.services.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/services/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.services.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.SearchEntities(tmpl), map[string]bool{"stock.services.search": true})).Methods(http.MethodGet)
	stck.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, entity.ViewEntity(tmpl), map[string]bool{"stock.services.view": true})).Methods(http.MethodGet)

	return r
}
