package handles

import (
	"fmt"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/stock/api"
	"github.com/louisevanderlith/stock/core"
	cars "github.com/louisevanderlith/vehicle/api"
	carcore "github.com/louisevanderlith/vehicle/core"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func GetStock(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Services", tmpl, "./views/categories.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllCategories(clnt, Endpoints["stock"], "A10")

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchStock(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Services", tmpl, "./views/categories.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllCategories(clnt, Endpoints["stock"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewStock(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Category View", tmpl, "./views/categoryview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		result := struct {
			Category core.Category
			Options  map[hsk.Key]string
		}{}

		clnt := CredConfig.Client(r.Context())
		cat, err := api.FetchCategory(clnt, Endpoints["stock"], key)

		if err != nil {
			log.Println("Fetch Service Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		result.Category = cat

		options, err := FetchCategoryOptions(strings.ToLower(cat.Name), clnt)

		if err != nil {
			log.Println("Fetch Options Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		result.Options = options

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func FetchCategoryOptions(name string, clnt *http.Client) (map[hsk.Key]string, error) {
	result := make(map[hsk.Key]string)
	switch name {
	case "cars":
		lst, err := cars.FetchAllVehicles(clnt, Endpoints["vehicle"], "A10")

		if err != nil {
			return nil, err
		}

		if lst.GetRecords() == nil {
			return result, nil
		}

		itor := lst.GetEnumerator()

		for itor.MoveNext() {
			curr := itor.Current().(hsk.Record)
			val := curr.GetValue().(carcore.Vehicle)
			result[curr.GetKey()] = fmt.Sprintf("%s %s %v", val.Series.Manufacturer, val.Series.Model, val.Series.Year)
		}
	case "clothing":
	case "spares":
	case "properties":
	case "utilities":
	case "tokens":
	}

	return result, nil
}
