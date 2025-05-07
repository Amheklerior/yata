package main

import (
	"fmt"
	"net/http"

	"github.com/Amheklerior/yata/server/internal/app"
	"github.com/Amheklerior/yata/server/internal/routes"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	// TODO: read from a .env file
	host, port := "localhost", 8080

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: routes.SetupRoutes(app),
	}

	app.Logger.Printf("YATA app server running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
