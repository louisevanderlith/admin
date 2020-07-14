package logbook

import (
	"github.com/louisevanderlith/admin/handles/menu"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func GetHistory(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "History", "./views/logbook/history.html")
	pge.AddMenu(menu.FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		err := ctx.Serve(http.StatusOK, pge.Page(nil, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
