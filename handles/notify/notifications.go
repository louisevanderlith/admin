package notify

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func GetNotifications(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Notifications", "./views/notify/notifications.html")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		err := ctx.Serve(http.StatusOK, pge.Page(nil, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
