package profile

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func Management(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Profile Manager", tmpl, "./views/profilemanager.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		pge.ChangeTitle("Profile & Security Manager")
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}