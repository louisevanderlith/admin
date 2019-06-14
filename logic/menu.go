package logic

import (
	"github.com/louisevanderlith/mango/control"
)

func GetMenu(path string) *control.Menu {
	return getItems(path)
}

func getItems(path string) *control.Menu {
	result := control.NewMenu(path)

	result.AddItem("", "#", "Artifact API", "fa-cloud", artifactChlidren(path))
	result.AddItem("", "#", "Blog API", "fa-pen", blogChildren(path))
	result.AddItem("", "#", "Comment API", "fa-chat", commentChildren(path))
	result.AddItem("", "#", "Comms API", "fa-fax", commsChildren(path))
	result.AddItem("", "#", "Entity API", "fa-person", entityChildren(path))
	result.AddItem("", "#", "Folio API", "fa-users", folioChildren(path))
	result.AddItem("", "#", "Funds API", "fa-money", fundsChildren(path))
	result.AddItem("", "#", "Game API", "fa-money", gameChildren(path))
	result.AddItem("", "#", "Logbook API", "fa-book", logbookChildren(path))
	result.AddItem("", "#", "Gate Proxy", "fa-globe", gateChildren(path))
	result.AddItem("", "#", "Notify API", "fa-chat", notifyChildren(path))
	result.AddItem("", "#", "Quote API", "fa-poes", quoteChildren(path))
	result.AddItem("", "#", "Router API", "fa-modem", routerChildren(path))
	result.AddItem("", "#", "Secure API", "fa-user-secret", secureChildren(path))

	result.AddItem("", "#", "Stock API", "fa-db", stockChildren(path))

	result.AddItem("", "#", "Theme API", "fa-brush", themeChildren(path))
	result.AddItem("", "#", "VIN API", "fa-notepad", vinChildren(path))
	result.AddItem("", "#", "Vehicle API", "fa-car", vehicleChildren(path))

	result.AddItem("", "#", "XChange API", "fa-x", xchangeChildren(path))

	return result
}

func artifactChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/uploads/A10", "Uploads", "fa-newspaper-o", nil)

	return children
}

func commentChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/comments/A10", "Comments", "fa-bath", nil)

	return children
}

func commsChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/comms/A10", "Messages", "fa-newspaper-o", nil)

	return children
}

func entityChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/entities/A10", "Entities", "fa-newspaper-o", nil)

	return children
}

func folioChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/profiles/A10", "Profiles", "fa-user", nil)

	return children
}

func fundsChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/accounts/A10", "Accounts", "fa-newspaper-o", nil)

	return children
}

func gameChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/heroes/A10", "Heroes", "fa-sword", nil)

	return children
}

func gateChildren(path string) *control.Menu {
	return control.NewMenu(path)
}

func logbookChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/history/A10", "Service History", "fa-blind", nil)

	return children
}

func notifyChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/notifications/A10", "Notifications", "fa-newspaper-o", nil)

	return children
}

func quoteChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/invoices/A10", "Invoices", "fa-newspaper-o", nil)

	return children
}

func stockChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/cars/A10", "Cars", "fa-newspaper-o", nil)
	children.AddItem("", "/parts/A10", "Parts", "fa-newspaper-o", nil)
	children.AddItem("", "/services/A10", "Services", "fa-newspaper-o", nil)

	return children
}

func routerChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/memory", "Memory", "fa-memory", nil)

	return children
}

func themeChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/templates", "Templates", "fa-camera", nil)
	children.AddItem("", "/css", "Stylesheets", "fa-blind", nil)

	return children
}

func secureChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/users/A10", "Users", "fa-user", nil)

	return children
}

func vehicleChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/vehicles/A10", "Vehicles", "fa-car", nil)

	return children
}

func xchangeChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/currency/A10", "Currency", "fa-money", nil)

	return children
}

func vinChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/vins/A10", "VIN", "fa-barcode", nil)

	return children
}

func blogChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("", "/blogs/A10", "Blogs", "fa-notepad", nil)

	return children
}
