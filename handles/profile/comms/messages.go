package comms

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/admin/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetMessages(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Messages", tmpl, "./views/comms/messages.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchComms("A10")

		if err != nil {
			log.Println("Fetch Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchMessages(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Messages", tmpl, "./views/comms/messages.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchComms(drx.FindParam(r, "pagesize"))

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

func ViewMessage(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Messages View", tmpl, "./views/comms/messageView.html")
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchCommsMessage(key.String())

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
