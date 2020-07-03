package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

func Index(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("index", mstr, tmpl)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		result := make(map[string]interface{})
		result["Menu"] = FullMenu()

		err := ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}

func Callback(w http.ResponseWriter, r *http.Request) {

}
