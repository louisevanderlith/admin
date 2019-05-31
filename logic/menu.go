package logic

import (
	"github.com/louisevanderlith/mango/control"
)

func GetMenu(path string) *control.Menu {
	return getItems(path)
}

func getItems(path string) *control.Menu {
	result := control.NewMenu(path)

	result.AddItem("#", "Artifact API", "fa-cloud", artifactChlidren(path))
	result.AddItem("#", "Comment API", "fa-chat", commentChildren(path))
	result.AddItem("#", "Comms API", "fa-fax", commsChildren(path))
	result.AddItem("#", "Folio API", "fa-users", folioChildren(path))
	result.AddItem("#", "Game API", "fa-money", gameChildren(path))
	result.AddItem("#", "Gate Proxy", "fa-globe", gateChildren(path))
	result.AddItem("#", "Logbook API", "fa-book", logbookChildren(path))
	result.AddItem("#", "Router API", "fa-modem", routerChildren(path))
	result.AddItem("#", "Theme API", "fa-brush", themeChildren(path))
	result.AddItem("#", "Secure API", "fa-user-secret", secureChildren(path))
	result.AddItem("#", "Vehicle API", "fa-car", vehicleChildren(path))
	result.AddItem("#", "VIN API", "fa-notepad", vinChildren(path))

	return result
}

func artifactChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/uploads/A10", "Uploads", "fa-newspaper-o", nil)

	return children
}

func commentChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/comments/A10", "Comments", "fa-bath", nil)

	return children
}

func commsChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/comms/A10", "Messages", "fa-newspaper-o", nil)

	return children
}

func folioChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/profiles/A10", "Profiles", "fa-user", nil)

	return children
}

func gameChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/heroes/A10", "Heroes", "fa-sword", nil)

	return children
}

func gateChildren(path string) *control.Menu {
	return control.NewMenu(path)
}

func logbookChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/history/A10", "Service History", "fa-blind", nil)

	return children
}

func routerChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/memory", "Memory", "fa-memory", nil)

	return children
}

func themeChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/templates", "Templates", "fa-camera", nil)
	children.AddItem("/css", "Stylesheets", "fa-blind", nil)

	return children
}

func secureChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/users/A10", "Users", "fa-user", nil)

	return children
}

func vehicleChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/vehicles/A10", "Vehicles", "fa-car", nil)

	return children
}

func vinChildren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/vins/A10", "VIN", "fa-barcode", nil)

	return children
}
