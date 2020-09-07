package artifact

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetUploads(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Uploads", tmpl, "./views/artifact/uploads.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchUploads("A10")

		if err != nil {
			log.Println("Fetch Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchUploads(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Uploads", tmpl, "./views/artifact/uploads.html")
	return func(w http.ResponseWriter, r *http.Request) {
		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchUploads(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewUpload(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Uploads View", tmpl, "./views/artifact/uploadsView.html")
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchUpload(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}
