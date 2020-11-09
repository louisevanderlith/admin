package main

import (
	"flag"
	"github.com/louisevanderlith/admin/handles"
	"net/http"
	"time"
)

func main() {
	host := flag.String("host", "http://127.0.0.1:8088", "This application's URL")
	clientId := flag.String("client", "admin", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	theme := flag.String("theme", "http://127.0.0.1:8093", "Theme URL")
	stock := flag.String("stock", "http://127.0.0.1:8101", "Stock URL")
	folio := flag.String("folio", "http://127.0.0.1:8090", "Folio URL")
	artifact := flag.String("artifact", "http://127.0.0.1:8082", "Artifact URL")
	funds := flag.String("funds", "http://127.0.0.1:8082", "Funds URL")
	vehicle := flag.String("vehicle", "http://127.0.0.1:8082", "Vehicle URL")
	game := flag.String("game", "http://127.0.0.1:8082", "Game URL")
	xchange := flag.String("xchange", "http://127.0.0.1:8088", "XChange URL")
	flag.Parse()

	ends := map[string]string{
		"issuer":   *issuer,
		"theme":    *theme,
		"stock":    *stock,
		"folio":    *folio,
		"artifact": *artifact,
		"funds":    *funds,
		"vehicle":  *vehicle,
		"game":     *game,
		"xchange":  *xchange,
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8088",
		Handler:      handles.SetupRoutes(*host, *clientId, *clientSecrt, ends),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
