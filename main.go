package main

/*
func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	profile := os.Getenv("PROFILE")
	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
	appName := os.Getenv("APPNAME")
	pubPath := path.Join(keyPath, pubName)

	// Register with router
	srv := bodies.NewService(appName, profile, pubPath, host, httpport, servicetype.APP)

	routr, err := do.GetServiceURL("", "Router.API", false)

	if err != nil {
		panic(err)
	}

	err = srv.Register(routr)

	if err != nil {
		panic(err)
	}

	err = droxolite.UpdateTheme(srv.ID)

	if err != nil {
		panic(err)
	}

	theme, err := element.GetDefaultTheme(host, srv.ID, profile)

	if err != nil {
		panic(err)
	}

	secur, err := do.GetServiceURL(srv.ID, "Auth.APP", true)

	if err != nil {
		panic(err)
	}

	err = theme.LoadTemplate("./views", "master.html")

	if err != nil {
		panic(err)
	}

	poxy := resins.NewColourEpoxy(srv, theme, secur, controllers.Index)
	routers.Setup(poxy)

	err = droxolite.Boot(poxy)

	if err != nil {
		panic(err)
	}
}
*/