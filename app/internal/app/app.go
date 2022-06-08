package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	config2 "go-bolvanka/internal/config"
	"go-bolvanka/internal/repository"
	"go-bolvanka/internal/service"
	v1 "go-bolvanka/internal/transport/http/v1"

	"go-bolvanka/pkg/httpserver"
	"syscall"

	"os"
	"os/signal"

	_ "go-bolvanka/docs"

	"go-bolvanka/pkg/logging"
	"go-bolvanka/pkg/postgres"
	"net/http"
)

type App struct {
	cfg        *config2.Config
	logger     *logging.Logger
	httpServer *http.Server
	pg         *postgres.Postgres

	// а
	itemService     *service.ItemService
	categoryService *service.CategoryService
}

func NewApp(config *config2.Config) (App, error) {
	logger := logging.New(config.Log.Level)
	logger.Info("router initializing")

	// Repository
	pg, err := postgres.New(config.PG.URL, postgres.MaxPoolSize(config.PG.PoolMax))
	if err != nil {
		logger.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	// сategory

	categoryRepo := repository.NewCategoryRepository(pg, logger)
	itemRepo := repository.NewItemRepository(pg, logger)

	return App{
		cfg:    config,
		logger: logger,
		pg:     pg,
		//TODO: придумать как сгруппировать чтобы не протаскивать сервисы через все методы
		itemService:     service.NewItemService(*itemRepo),
		categoryService: service.NewCategoryService(*categoryRepo),
	}, nil
}

func (a *App) Run() {
	err := a.startHTTP()
	if err != nil {
		a.logger.Fatal("Failed app configuration", err)
	}
}

func (a *App) startHTTP() error {
	a.logger.Info("start HTTP")
	// Init gin handler

	a.logger.Info(fmt.Sprintf("bind application to host: %s", a.cfg.HTTP))

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, a.logger, *a.categoryService, *a.itemService)
	httpServer := httpserver.New(handler, httpserver.Port(a.cfg.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
	select {
	case s := <-interrupt:
		a.logger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		a.logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		a.logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
	return nil
}
