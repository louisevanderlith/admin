package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

func Index(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		result := make(map[string]string)
		result["Comms.API"] = "/comms/messages/A10"
		result["Router.API"] = "/router/memory"
		result["Folio.API"] = "/folio/profiles/A10"

		err := ctx.Serve(http.StatusOK, mix.Page("Index", result, ctx.GetTokenInfo(), mstr, tmpl))

		if err != nil {
			log.Println(err)
		}
	}
}

func Callback(w http.ResponseWriter, r *http.Request) {

}
