package app

import (
	"database/sql"
	"fmt"
	"github/SpaceBuckett/CarGo/Internal/api"
	"github/SpaceBuckett/CarGo/Internal/store"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	PayloadHandler *api.PayloadHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, fmt.Errorf("error store.open: %w", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	payloadHandler := api.NewPaylaodHandler()

	app := &Application{
		Logger:         logger,
		PayloadHandler: payloadHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "STATUS GREEN\n")
}
