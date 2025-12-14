package app

import (
	"fmt"
	"github/SpaceBuckett/CarGo/Internal/api"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	PayloadHandler *api.PayloadHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// OUR STORES GO HERE

	// OUR ROUTES GO HERE
	payloadHandler := api.NewPaylaodHandler()

	app := &Application{
		Logger:         logger,
		PayloadHandler: payloadHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "STATUS GREEN\n")
}
