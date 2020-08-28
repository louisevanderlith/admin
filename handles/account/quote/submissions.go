package quote

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func GetSubmissions(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Submissions", tmpl, "./views/quote/submissions.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
