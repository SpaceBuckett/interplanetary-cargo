package main

import (
	"flag"
	"fmt"
	"github/SpaceBuckett/CarGo/Internal/app"
	"github/SpaceBuckett/CarGo/Internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "SERVER PORT")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Welcome to CarGo!")

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Listening on port %d", port)
	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}

}
