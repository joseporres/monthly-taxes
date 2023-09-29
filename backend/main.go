package main

import (
	"context"
	"fmt"

	"backend/config"
	"backend/database"
	"backend/global"
	"backend/internal/api"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/settings"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	config.LoadEnvFile()
	fmt.Println(global.TOKEN)
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
