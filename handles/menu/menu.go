package menu

import "github.com/louisevanderlith/droxolite/menu"

func FullMenu() *menu.Menu {
	m := menu.NewMenu()
	m.AddItem(artifactMenu())
	m.AddItem(blogMenu())
	m.AddItem(commentMenu())
	m.AddItem(commsMenu())
	m.AddItem(curityMenu())
	m.AddItem(entityMenu())
	m.AddItem(fundsMenu())
	m.AddItem(gameMenu())
	m.AddItem(logbookMenu())
	m.AddItem(notifyMenu())
	m.AddItem(quoteMenu())
	m.AddItem(stockMenu())
	m.AddItem(themeMenu())
	m.AddItem(vehicleMenu())
	m.AddItem(vinMenu())
	m.AddItem(xchangeMenu())

	return m
}

func artifactMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("uploads", "/artifact/uploads", "Uploads", nil),
	}

	return menu.NewItem("artifact", "#", "Artifacts", chldrn)
}

func blogMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("articles", "/blog/articles", "Articles", nil),
	}

	return menu.NewItem("blog", "#", "Blog", chldrn)
}

func commentMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("comments", "/comment/messages", "Messages", nil),
	}

	return menu.NewItem("comment", "#", "Comments", chldrn)
}

func commsMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("messages", "/comms/messages", "Messages", nil),
	}

	return menu.NewItem("comms", "#", "Comms", chldrn)
}

func curityMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("analytics", "/curity/analytics", "Analytics", nil),
		menu.NewItem("profiles", "/curity/profiles", "Profiles", nil),
		menu.NewItem("resources", "/curity/resources", "Resources", nil),
		menu.NewItem("users", "/curity/users", "Users", nil),
	}

	return menu.NewItem("curity", "#", "Security", chldrn)
}

func entityMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("entities", "/entity/entities", "Entities", nil),
	}

	return menu.NewItem("entity", "#", "Entity", chldrn)
}

func fundsMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("accounts", "/funds/accounts", "Accounts", nil),
	}

	return menu.NewItem("funds", "#", "Funds", chldrn)
}

func gameMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("heroes", "/game/heroes", "Heroes", nil),
	}

	return menu.NewItem("game", "#", "Game", chldrn)
}

func logbookMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("history", "/logbook/history", "History", nil),
	}

	return menu.NewItem("logbook", "#", "Logbook", chldrn)
}

func notifyMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("notifications", "/notify/notifications", "Notifications", nil),
	}

	return menu.NewItem("notify", "#", "Notify", chldrn)
}

func quoteMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("submissions", "/quote/submissions", "Submissions", nil),
	}

	return menu.NewItem("quote", "#", "Quote", chldrn)
}

func stockMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("cars", "/stock/cars", "Cars", nil),
		menu.NewItem("parts", "/stock/parts", "Parts", nil),
		menu.NewItem("services", "/stock/services", "Services", nil),
		menu.NewItem("properties", "/stock/properties", "Properties", nil),
	}

	return menu.NewItem("stock", "#", "Stock", chldrn)
}

func themeMenu() menu.Item {
	//TODO move to 'profile' section
	chldrn := []menu.Item{
		menu.NewItem("assets", "/theme/assets", "Assets", nil),
		menu.NewItem("stylesheets", "/theme/stylesheets", "Stylesheets", nil),
		menu.NewItem("templates", "/theme/templates", "Templates", nil),
	}

	return menu.NewItem("theme", "#", "Theme", chldrn)
}

func vehicleMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("vehicles", "/vehicle/vehicles", "Vehicles", nil),
	}

	return menu.NewItem("vehicle", "#", "Vehicle", chldrn)
}

func vinMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("vins", "/vin/vin", "VIN", nil),
		menu.NewItem("regions", "/vin/regions", "Regions", nil),
	}

	return menu.NewItem("vin", "#", "VIN", chldrn)
}

func xchangeMenu() menu.Item {
	chldrn := []menu.Item{
		menu.NewItem("credits", "/xchange/credits", "Credits", nil),
	}

	return menu.NewItem("xchange", "#", "XChange", chldrn)
}
