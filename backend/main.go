package main

import (
	"context"
	"fmt"

	"backend/database"
	"backend/internal/api"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/settings"

	"backend/config"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),

		fx.Invoke(
			setLifeCycle,
		),
	)

	config.LoadEnvFile()
	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
