package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ger9000/memes-as-a-service/internal/delivery/rest"
	"github.com/ger9000/memes-as-a-service/internal/delivery/rest/billing"
	"github.com/ger9000/memes-as-a-service/internal/delivery/rest/healthcheck"
	"github.com/ger9000/memes-as-a-service/internal/delivery/rest/memes"
	"github.com/ger9000/memes-as-a-service/internal/infrastructure"
	"github.com/ger9000/memes-as-a-service/internal/shared/config"
	"github.com/ger9000/memes-as-a-service/internal/shared/datasource"
	"github.com/ger9000/memes-as-a-service/internal/shared/http/middleware"
	billingActions "github.com/ger9000/memes-as-a-service/internal/usecases/billing"
	memesAction "github.com/ger9000/memes-as-a-service/internal/usecases/memes"
	"github.com/rs/zerolog/log"
)

type App struct {
	server *http.Server
}

func InitApp() (*App, error) {
	routerHandlers := initializeDependencies()

	mux := rest.NewRouter(routerHandlers)
	server := NewServer(mux)
	return &App{
		server,
	}, nil
}

func (app *App) Start() {
	go func() {
		log.Info().Msgf("App MaaS(memes as a service) at port %d started!", config.GetInstance().Server.Port)
		if err := app.server.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener
			log.Fatal().Err(err).Msgf("HTTP server ListenAndServe fail")
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	gracefullyCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := app.server.Shutdown(gracefullyCtx); err != nil {
		log.Fatal().Err(err).Msgf("server shutdown fail")
	}

	log.Info().Msg("server has shutdown")
}

func initializeDependencies() rest.RouterHandlers {
	getAllAction := memesAction.NewGetAllAction()
	db, err := datasource.New()
	if err != nil {
		panic(err)
	}
	findTrackAction := billingActions.NewFindTrackAction(infrastructure.NewRepository(db))
	consumeAvailableAction := billingActions.NewConsumeAvailableCallAction(infrastructure.NewRepository(db))
	rechargeAvailablesAction := billingActions.NewRechargeAvailableCallAction(infrastructure.NewRepository(db))

	return rest.RouterHandlers{
		HealthCheckController:           healthcheck.NewController(),
		GetAllMemesController:           memes.NewGetAllController(getAllAction),
		RechargeAvailableCallController: billing.NewRechargeAvailableCallController(rechargeAvailablesAction),
		APICallsTrackerMiddleware:       *middleware.NewCallsTracker(findTrackAction, consumeAvailableAction),
	}
}
