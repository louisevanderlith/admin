package main

import (
	"flag"
	"github.com/louisevanderlith/admin/handles"
	"github.com/louisevanderlith/kong"
	"net/http"
	"time"

	"github.com/louisevanderlith/droxolite"
)

func main() {
	clientId := flag.String("client", "mango.admin", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	authrty := flag.String("authority", "http://localhost:8094", "Authority Provider's URL")
	securty := flag.String("security", "http://localhost:8086", "Security Provider's URL")

	flag.Parse()

	tkn, err := kong.FetchToken(http.DefaultClient, *securty, *clientId, *clientSecrt, "theme.assets.download", "theme.assets.view")

	if err != nil {
		panic(err)
	}

	clms, err := kong.Exchange(http.DefaultClient, tkn, *clientId, *clientSecrt, *securty+"/info")

	if err != nil {
		panic(err)
	}

	err = droxolite.UpdateTemplate(tkn, clms)

	if err != nil {
		panic(err)
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8088",
		Handler:      handles.SetupRoutes(*clientId, *clientSecrt, *securty, *authrty),
	}

	err = srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
