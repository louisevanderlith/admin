package notify

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func GetNotifications(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		err := ctx.Serve(http.StatusOK, mix.Page("notifications", nil, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}
