package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/admin/handles/artifact"
	"github.com/louisevanderlith/admin/handles/blog"
	"github.com/louisevanderlith/admin/handles/comment"
	"github.com/louisevanderlith/admin/handles/comms"
	"github.com/louisevanderlith/admin/handles/curity"
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

	crty := r.PathPrefix("/curity").Subrouter()
	crty.HandleFunc("/profiles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.GetProfiles(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchProfiles(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchProfiles(tmpl), "secure.profile.search")).Methods(http.MethodGet)
	crty.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.ViewProfile(tmpl), "secure.profile.view")).Methods(http.MethodGet)

	crty.HandleFunc("/resources", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.GetResource(tmpl), "secure.resource.search")).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchResource(tmpl), "secure.resource.search")).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.SearchResource(tmpl), "secure.resource.search")).Methods(http.MethodGet)
	crty.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, curity.ViewResource(tmpl), "secure.resource.view")).Methods(http.MethodGet)

	return r
}
