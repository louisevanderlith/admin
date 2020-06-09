package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/admin/handles/artifact"
	"github.com/louisevanderlith/admin/handles/blog"
	"github.com/louisevanderlith/admin/handles/comment"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/kong"
	"net/http"
)

func SetupRoutes(clnt, scrt, secureUrl, authUrl string) http.Handler {
	mstr, tmpl, err := droxolite.LoadTemplate("./views", "master.html")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", Index(mstr, tmpl))

	artfct := r.PathPrefix("/artifact").Subrouter()
	artfct.HandleFunc("/uploads", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, artifact.GetUploads(mstr, tmpl), "artifact.uploads.view")).Methods(http.MethodGet)

	blg := r.PathPrefix("/blog").Subrouter()
	blg.HandleFunc("/articles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, blog.GetArticles(mstr, tmpl), "blog.articles.view")).Methods(http.MethodGet)

	cmmnt := r.PathPrefix("/comment").Subrouter()
	cmmnt.HandleFunc("/messages", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, comment.GetMessages(mstr, tmpl), "artifact.uploads.view")).Methods(http.MethodGet)

	cmms := r.PathPrefix("/comms").Subrouter()
	cmms.HandleFunc("/messages", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, artifact.GetUploads(mstr, tmpl), "artifact.uploads.view")).Methods(http.MethodGet)

	return r
}
