package handles

import (
	"fmt"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	house "github.com/louisevanderlith/house/api"
	housecore "github.com/louisevanderlith/house/core"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	parts "github.com/louisevanderlith/parts/api"
	partscore "github.com/louisevanderlith/parts/core"
	"github.com/louisevanderlith/stock/api"
	"github.com/louisevanderlith/stock/core"
	"github.com/louisevanderlith/stock/core/categories"
	utility "github.com/louisevanderlith/utility/api"
	utilitycore "github.com/louisevanderlith/utility/core"
	cars "github.com/louisevanderlith/vehicle/api"
	carcore "github.com/louisevanderlith/vehicle/core"
	wear "github.com/louisevanderlith/wear/api"
	wearcore "github.com/louisevanderlith/wear/core"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func GetStockCategories(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Categories", tmpl, "./views/categories.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchAllCategories(clnt, Endpoints["stock"], "A10")

		if err != nil {
			log.Println("Fetch Categories Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		pge.ChangeTitle("Stock Categories")
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchStockCategories(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Categories", tmpl, "./views/categories.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchAllCategories(clnt, Endpoints["stock"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Categories Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		pge.ChangeTitle("Stock Categories")
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewStockCategory(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Category View", tmpl, "./views/categoryview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		result := struct {
			Category   core.Category
			Options    map[hsk.Key]string
			CreatePath string
		}{}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		cat, err := api.FetchCategory(clnt, Endpoints["stock"], key)

		if err != nil {
			log.Println("Fetch Service Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		catBase := categories.StringEnum(cat.BaseCategory)
		result.Category = cat
		result.CreatePath = fmt.Sprintf("/%s/create?from=%s", strings.ToLower(catBase), key.String())

		options, err := FetchCategoryOptions(catBase, clnt)

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
	actions := map[string]func(client *http.Client) (map[hsk.Key]string, error){
		"Cars":       fetchVehicleOptions,
		"Clothing":   fetchClothingOptions,
		"Spares":     fetchPartOptions,
		"Properties": fetchPropertyOptions,
		"Utilities":  fetchUtilityOptions,
		"Tokens":     fetchTokenOptions,
	}

	act, ok := actions[name]

	if !ok {
		return nil, fmt.Errorf("unknown option: %s", name)
	}

	return act(clnt)
}

func fetchVehicleOptions(clnt *http.Client) (map[hsk.Key]string, error) {
	result := make(map[hsk.Key]string)

	lst, err := cars.FetchAllVehicles(clnt, Endpoints["vehicle"], "A10")

	if err != nil {
		return nil, err
	}

	if lst.GetRecords() == nil || !lst.Any() {
		return result, nil
	}

	itor := lst.GetEnumerator()

	for itor.MoveNext() {
		curr := itor.Current().(hsk.Record)
		val := curr.GetValue().(carcore.Vehicle)
		result[curr.GetKey()] = fmt.Sprintf("%s %s %v", val.Series.Manufacturer, val.Series.Model, val.Series.Year)
	}

	return result, nil
}

func fetchUtilityOptions(clnt *http.Client) (map[hsk.Key]string, error) {
	result := make(map[hsk.Key]string)

	lst, err := utility.FetchAllServices(clnt, Endpoints["utility"], "A10")

	if err != nil {
		return nil, err
	}

	if lst.GetRecords() == nil || !lst.Any() {
		return result, nil
	}

	itor := lst.GetEnumerator()

	for itor.MoveNext() {
		curr := itor.Current().(hsk.Record)
		val := curr.GetValue().(*utilitycore.Service)
		result[curr.GetKey()] = fmt.Sprintf("%s @ %s", val.Description, val.Location)
	}

	return result, nil
}

func fetchTokenOptions(clnt *http.Client) (map[hsk.Key]string, error) {
	result := make(map[hsk.Key]string)

	/*lst, err := xchange.FetchAllTokens(clnt, Endpoints["xchange"], "A10")

	if err != nil {
		return nil, err
	}

	if !lst.Any() {
		return result, nil
	}

	itor := lst.GetEnumerator()

	for itor.MoveNext() {
		curr := itor.Current().(hsk.Record)
		val := curr.GetValue().(xchangecore.Token)
		result[curr.GetKey()] = fmt.Sprintf("%s %s %v", val.Series.Manufacturer, val.Series.Model, val.Series.Year)
	}*/

	return result, nil
}

func fetchClothingOptions(clnt *http.Client) (map[hsk.Key]string, error) {
	result := make(map[hsk.Key]string)

	lst, err := wear.FetchAllClothing(clnt, Endpoints["wear"], "A10")

	if err != nil {
		return nil, err
	}

	if lst.GetRecords() == nil || !lst.Any() {
		return result, nil
	}

	itor := lst.GetEnumerator()

	for itor.MoveNext() {
		curr := itor.Current().(hsk.Record)
		val := curr.GetValue().(wearcore.Clothing)
		result[curr.GetKey()] = fmt.Sprintf("%s %s %s", val.Brand, val.Type, val.Colour)
	}

	return result, nil
}

func fetchPropertyOptions(clnt *http.Client) (map[hsk.Key]string, error) {
	result := make(map[hsk.Key]string)

	lst, err := house.FetchAllProperties(clnt, Endpoints["house"], "A10")

	if err != nil {
		return nil, err
	}

	if lst.GetRecords() == nil || !lst.Any() {
		return result, nil
	}

	itor := lst.GetEnumerator()

	for itor.MoveNext() {
		curr := itor.Current().(hsk.Record)
		val := curr.GetValue().(housecore.Property)
		result[curr.GetKey()] = fmt.Sprintf("%s %s", val.Type, val.Address)
	}

	return result, nil
}

func fetchPartOptions(clnt *http.Client) (map[hsk.Key]string, error) {
	result := make(map[hsk.Key]string)

	lst, err := parts.FetchAllSpares(clnt, Endpoints["parts"], "A10")

	if err != nil {
		return nil, err
	}

	if lst.GetRecords() == nil || !lst.Any() {
		return result, nil
	}

	itor := lst.GetEnumerator()

	for itor.MoveNext() {
		curr := itor.Current().(hsk.Record)
		val := curr.GetValue().(partscore.Spare)
		result[curr.GetKey()] = val.Number
	}

	return result, nil
}
